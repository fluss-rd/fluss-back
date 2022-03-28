package repository

import (
	"context"
	"fmt"
	"sort"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/flussrd/fluss-back/app/reporting/models"
	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
	calculator "github.com/flussrd/fluss-back/app/shared/wqi-calculator"
)

const (
	organizationName = "fluss"
	bucketName       = "fluss"
)

type influxRepository struct {
	client influxdb2.Client
}

func New(client influxdb2.Client) repository.Repository {
	return influxRepository{
		client: client,
	}
}

func buildGetDataByModuleQuery(moduleID string, options models.SearchOptions) string {
	// query := `from(bucket: "modules-data")
	// |> range(start: -48h, stop: now())
	// |> filter(fn: (r) => r["_measurement"] == "water-sensor")
	// |> filter(fn: (r) => r["moduleID"] == "MDL8dab9bcded8b4a0a9b18a9b8e2e0c758")
	// |> aggregateWindow(every: 1h, fn: mean, createEmpty: false)
	// |> yield(name: "mean")`

	query := fmt.Sprintf(`from(bucket: "%s")`, bucketName)

	start := "-48h" // we're defaulting the last 48 hours
	stop := "now()"

	if !options.Start.IsZero() {
		start = fmt.Sprint(options.Start.Format(time.RFC3339))
	}

	if !options.End.IsZero() {
		stop = fmt.Sprint(options.End.Format(time.RFC3339))
	}

	query += fmt.Sprintf(`|> range(start: %s, stop: %s)`, start, stop)

	query += `|> filter(fn: (r) => r["_measurement"] == "water-sensor")`

	query += fmt.Sprintf(`|> filter(fn: (r) => r["moduleID"] == "%s")`, moduleID)

	if options.Cardinality != "" && options.AggregationWindow != "" {
		query += fmt.Sprintf(`|> aggregateWindow(every: %s, fn: %s, createEmpty: false)`, options.Cardinality, options.AggregationWindow)
		query += fmt.Sprintf(`|> yield(name: "%s")`, options.AggregationWindow)
	}

	return query
}

func (repo influxRepository) GetDataByModule(ctx context.Context, moduleID string, searchOptions models.SearchOptions) (models.Report, error) {
	queryAPI := repo.client.QueryAPI(organizationName)

	query := buildGetDataByModuleQuery(moduleID, searchOptions)

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return models.Report{}, err
	}

	var report *models.Report
	dataPerModuleAndTime := map[string]*models.Data{}
	locationPerModuleAndTime := map[string]map[string]float64{}
	parameters := map[string][]models.Parameter{}

	for result.Next() {
		riverID, _ := result.Record().ValueByKey("riverID").(string)
		if riverID == "" {
			continue
		}

		if report == nil {
			riverID, _ := result.Record().ValueByKey("riverID").(string)
			moduleID, _ := result.Record().ValueByKey("moduleID").(string)
			report = &models.Report{
				ModuleID: moduleID,
				RiverID:  riverID,
			}
		}

		key := fmt.Sprintf("%v:%s", result.Record().ValueByKey("moduleID"), result.Record().Time().String())
		if dataPerModuleAndTime[key] == nil {
			dataPerModuleAndTime[key] = &models.Data{
				LastDate: result.Record().Time(),
			}
		}

		if result.Record().Time().After(dataPerModuleAndTime[key].LastDate) {
			dataPerModuleAndTime[key].LastDate = result.Record().Time()
		}

		// We're grouping locations by data groups
		if locationPerModuleAndTime[key] == nil {
			locationPerModuleAndTime[key] = map[string]float64{}
		}

		// THIS has go AFTER the map is initialized in last step
		if result.Record().Field() == "lat" {
			lat, _ := result.Record().Value().(float64)
			locationPerModuleAndTime[key]["lat"] = lat
			continue
		}

		if result.Record().Field() == "lng" {
			lng, _ := result.Record().Value().(float64)
			locationPerModuleAndTime[key]["lng"] = lng
			continue
		}

		if parameters[key] == nil {
			parameters[key] = []models.Parameter{}
		}

		parameters[key] = append(parameters[key], models.Parameter{
			Parameter: calculator.Parameter{
				Name:  calculator.ParameterType(result.Record().Field()),
				Value: result.Record().Value().(float64)},
			Date: result.Record().Time(),
		})
	}

	if len(parameters) == 0 {
		return models.Report{}, nil
	}

	// Assigning parameters and location to data
	for k := range dataPerModuleAndTime {
		dataPerModuleAndTime[k].Parameters = parameters[k]
		dataPerModuleAndTime[k].Location = models.Location{
			Latitude:  locationPerModuleAndTime[k]["lat"],
			Longitude: locationPerModuleAndTime[k]["lat"],
		}
	}

	datas := []models.Data{}

	// Assigning the data the report
	for _, v := range dataPerModuleAndTime {
		datas = append(datas, *v)
	}

	// Ugly, i know, but the parsing technique with the maps screws the order up :(
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].LastDate.Before(datas[j].LastDate)
	})

	report.Data = datas

	return *report, nil
}

func buildGetAllModulesSummaryQuery(options models.SearchOptions) string {
	query := fmt.Sprintf(`from(bucket: "%s")`, bucketName)
	query += "|> range(start: -24h, stop: now())" // we're making this a day cause we want the last one for each module in the last day, just to get one. Is this business logic?
	query += `|> filter(fn: (r) => r["_measurement"] == "water-sensor")`

	if options.RiverID != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r["riverID"] == "%s")`, options.RiverID)
	}

	query += `|> last()`

	return query
}

func (repo influxRepository) GetAllModulesSummary(ctx context.Context, options models.SearchOptions) ([]models.Report, error) {
	queryAPI := repo.client.QueryAPI(organizationName)

	query := buildGetAllModulesSummaryQuery(options)

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	reportsPerModule := map[string]*models.Report{} // oinly one cause we're only getting the last one

	lastDatePerModule := map[string]time.Time{}

	for result.Next() {
		// TODO: add error handling for the case when !ok
		moduleID, _ := result.Record().ValueByKey("moduleID").(string)
		riverID, _ := result.Record().ValueByKey("riverID").(string)

		// This hast to go first to ensure report is not nil and the data array is not empty
		if reportsPerModule[moduleID] == nil {
			reportsPerModule[moduleID] = &models.Report{
				ModuleID: moduleID,
				RiverID:  riverID,
				Data:     []models.Data{{}}, // Only one, just the last one
			}
		}

		// if lastDatePerModule[moduleID].IsZero() {
		// 	lastDatePerModule[moduleID] = result.Record().Time()
		// }

		if result.Record().Time().After(lastDatePerModule[moduleID]) {
			lastDatePerModule[moduleID] = result.Record().Time()
		}

		if result.Record().Field() == "lat" {
			lat, _ := result.Record().Value().(float64)
			reportsPerModule[moduleID].Data[0].Location.Latitude = lat
			continue
		}

		if result.Record().Field() == "lng" {
			lng, _ := result.Record().Value().(float64)
			reportsPerModule[moduleID].Data[0].Location.Longitude = lng
			continue
		}

		if calculator.IsValidParamType(calculator.ParameterType(result.Record().Field())) {
			reportsPerModule[moduleID].Data[0].Parameters = append(reportsPerModule[moduleID].Data[0].Parameters, models.Parameter{
				Parameter: calculator.Parameter{
					Name:  calculator.ParameterType(result.Record().Field()),
					Value: result.Record().Value().(float64), // TODO: error handling
				},
				Date: result.Record().Time(),
			})
		}
	}

	for key := range reportsPerModule {
		reportsPerModule[key].LastUpdated = lastDatePerModule[key]
		for index, _ := range reportsPerModule[key].Data {
			reportsPerModule[key].Data[index].LastDate = lastDatePerModule[key]
		}
	}

	output := make([]models.Report, len(reportsPerModule))

	index := 0
	for _, v := range reportsPerModule {
		output[index] = *v
		index++
	}

	return output, nil
}

func buildGetRiverSummaryQuery(riverID string) string {
	query := fmt.Sprintf(`times = from(bucket: "modules-data")
    |> range(start: -24h, stop: now())
    |> filter(fn: (r) => r["_measurement"] == "water-sensor" and r["riverID"] == "%s")
    |> last()
    |> group(columns: ["moduleID"], mode:"by")
    |> group(columns: ["_field"], mode:"by")
    |> max(column: "_time")
	|> keep(columns: ["_time", "_field"])`, riverID)

	query += fmt.Sprintf(`averages = from(bucket: "modules-data")
	|> range(start: 0, stop: now())
	|> filter(fn: (r) => r["_measurement"] == "water-sensor" and r["riverID"] == "%s")
	|> last()
	|> group(columns: ["moduleID"], mode:"by")
	|> group(columns: ["_field"], mode:"by")
	|> mean()`, riverID)

	query += `join(
	tables: {times:times, averages:averages},
	on: ["_field"]
  )`

	return query
}

func (repo influxRepository) GetRiverSummary(ctx context.Context, riverID string) (models.Report, error) {
	queryAPI := repo.client.QueryAPI(organizationName)

	query := buildGetRiverSummaryQuery(riverID)

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return models.Report{}, nil
	}

	report := models.Report{RiverID: riverID, Data: []models.Data{{}}}
	lastDate := time.Time{}

	for result.Next() {
		if report.LastUpdated.Before(result.Record().Time()) {
			report.LastUpdated = result.Record().Time()
		}

		if calculator.IsValidParamType(calculator.ParameterType(result.Record().Field())) {
			report.Data[0].Parameters = append(report.Data[0].Parameters, models.Parameter{
				Parameter: calculator.Parameter{
					Name:  calculator.ParameterType(result.Record().Field()),
					Value: result.Record().Value().(float64), // TODO: error handling
				},
				Date: result.Record().Time(),
			})
		}

		if result.Record().Time().After(lastDate) {
			lastDate = result.Record().Time()
		}
	}

	time.Now().UnixNano()
	report.Data[0].LastDate = lastDate

	return report, nil
}
