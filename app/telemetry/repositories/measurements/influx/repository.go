package repository

import (
	"context"

	"github.com/flussrd/fluss-back/app/telemetry/models"
	repository "github.com/flussrd/fluss-back/app/telemetry/repositories/measurements"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

func (repo influxRepository) SaveMeasurement(ctx context.Context, message models.Message) error {
	writeApi := repo.client.WriteAPIBlocking(organizationName, bucketName)

	// TODO: we should add the riverID as a tag for being able to query by river
	point := influxdb2.NewPointWithMeasurement("water-sensor").AddTag("moduleID", message.ModuleID).AddTag("riverID", message.RiverID)

	for _, measurement := range message.Measurements {
		point = point.AddField(string(measurement.Name), measurement.Value)
	}

	point = point.SetTime(message.Date)

	// TODO: handle errors coming from the influx
	return writeApi.WritePoint(ctx, point)
}
