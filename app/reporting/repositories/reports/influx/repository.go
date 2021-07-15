package repository

import (
	"context"
	"fmt"
	"sort"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/flussrd/fluss-back/app/reporting/models"
	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
)

const (
	organizationName = "Fluss"
	bucketName       = "modules-data"
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

	fmt.Println(options.Start.IsZero())
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

	fmt.Println(query)

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

		if locationPerModuleAndTime[key] == nil {
			locationPerModuleAndTime[key] = map[string]float64{}
		}

		// THIS has go AFTER the map is initialized in last step
		if result.Record().Field() == "lat" {
			lat, _ := result.Record().Value().(float64)
			locationPerModuleAndTime[key]["lat"] = lat
		}

		if result.Record().Field() == "lng" {
			lng, _ := result.Record().Value().(float64)
			locationPerModuleAndTime[key]["lng"] = lng
		}

		if parameters[key] == nil {
			parameters[key] = []models.Parameter{}
		}

		parameters[key] = append(parameters[key], models.Parameter{
			Name:  result.Record().Field(),
			Value: result.Record().Value().(float64),
			Date:  result.Record().Time(),
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
