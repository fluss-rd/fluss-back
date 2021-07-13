package repository

import (
	"context"
	"testing"
	"time"

	"github.com/flussrd/fluss-back/app/telemetry/models"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	c := require.New(t)

	client := influxdb2.NewClient("http://localhost:8086", "sOR6NRnbCE25039soFTiK468obOjXxd-yKPpvBmZ4eCwdpDFAQ0pey7ifhySeZoYWZmlAPmqIg9nDerxzgqg1w==")

	repo := New(client)
	c.NotNil(repo)

	message := models.Message{
		ModuleID: "MDL-TEST",
		Date:     time.Now(),
		Measurements: []models.Measurement{
			models.Measurement{
				Name:  "temperature",
				Value: 5.02,
			},
			{
				Name:  "ph",
				Value: 5,
			},
		},
	}

	err := repo.SaveMeasurement(context.Background(), message)
	c.Nil(err)
}
