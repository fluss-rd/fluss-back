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
	err := repo.GetData(context.Background())
	c.Nil(err)
}
