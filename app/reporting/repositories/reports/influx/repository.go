package repository

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/flussrd/fluss-back/app/reporting/models"
	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
)

const (
	organizationName = "Fluss"
)

type influxRepository struct {
	client influxdb2.Client
}

func New(client influxdb2.Client) repository.Repository {
	return influxRepository{
		client: client,
	}
}

func (repo influxRepository) GetDataByModule(ctx context.Context, moduleID string) (models.Report, error) {
	// URGENT TODO: add the wqi to the measurement
	queryAPI := repo.client.QueryAPI(organizationName)

	// query := `from(bucket: "modules-data")
	// |> range(start: -48h, stop: now())
	// |> filter(fn: (r) => r["_measurement"] == "water-sensor")
	// |> aggregateWindow(every: 1h, fn: mean, createEmpty: false)
	// |> yield(name: "mean")`

	query := `from(bucket: "modules-data")
	|> range(start: -48h, stop: now())
	|> filter(fn: (r) => r["_measurement"] == "water-sensor")
	|> filter(fn: (r) => r["moduleID"] == "MDL8dab9bcded8b4a0a9b18a9b8e2e0c758")
	|> aggregateWindow(every: 1h, fn: mean, createEmpty: false)
	|> yield(name: "mean")`

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return models.Report{}, err
	}

	var report *models.Report
	dataPerModuleAndTime := map[string]*models.Data{}
	locationPerModuleAndTime := map[string]map[string]float64{}
	parameters := map[string][]models.Parameter{}

	for result.Next() {
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
				Date: result.Record().Time(),
			}
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
		})
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

	report.Data = datas

	return *report, nil
}
