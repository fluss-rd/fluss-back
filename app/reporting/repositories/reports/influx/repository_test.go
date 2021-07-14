package repository

import (
	"context"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/stretchr/testify/require"
)

func TestGetData(t *testing.T) {
	c := require.New(t)

	client := influxdb2.NewClient("http://localhost:8086", "e_6ya-hyjJeZHtmJd0IvX3f-2Z39Eab_KGKw95-tjii0Az8SrojeS8W2KYoDmW1xUMYc42ocJav6AkuOrb84jQ==")
	repo := New(client)
	report, err := repo.GetDataByModule(context.Background(), "MDL8dab9bcded8b4a0a9b18a9b8e2e0c758")
	c.Nil(err)
	c.Equal("MDL8dab9bcded8b4a0a9b18a9b8e2e0c758", report.ModuleID)
	c.NotEmpty(report.Data)
	c.NotEmpty(report.Data[0].Parameters)
}
