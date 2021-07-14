package repository

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
)

// change this pls
type value struct {
	ph          float64
	temperature float64
}

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

func (repo influxRepository) GetData(ctx context.Context) error {
	// URGENT TODO: add the wqi to the measurement
	queryAPI := repo.client.QueryAPI(organizationName)

	query := `from(bucket: "modules-data")
	|> range(start: -48h, stop: now())
	|> filter(fn: (r) => r["_measurement"] == "water-sensor")
	|> aggregateWindow(every: 1h, fn: mean, createEmpty: false)
	|> yield(name: "mean")`

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		fmt.Println("Hola")
		return err
	}

	values := map[string]*value{}

	for result.Next() {
		key := fmt.Sprintf("%v:%s", result.Record().ValueByKey("moduleID"), result.Record().Time().String())
		if values[key] == nil {
			values[key] = &value{}
		}

		switch result.Record().Field() {
		case "ph":
			values[key].ph = result.Record().Value().(float64)
		case "tmp":
			values[key].temperature = result.Record().Value().(float64)
		}
	}

	list := make([]value, len(values))

	index := 0
	for _, v := range values {
		list[index] = *v
		index++
	}

	fmt.Println(list[0])

	return nil
}
