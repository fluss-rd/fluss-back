package service

import (
	"sort"
	"testing"

	"github.com/flussrd/fluss-back/app/telemetry/models"
	"github.com/stretchr/testify/require"
)

func TestService_GetMeasurementsFromMessageBody(t *testing.T) {
	c := require.New(t)

	service := service{}

	body := `TEMP*27.19;pH?5.85/TDS+0.00)TURB!48.22(D.O%1729=`

	measurements, err := service.getMeasurementsFromMessageBody(body)
	c.Nil(err)
	c.Len(measurements, 5)

	sort.Slice(measurements, func(i, j int) bool {
		return measurements[i].Name < measurements[j].Name
	})

	c.Equal(float64(1729), measurements[0].Value)
	c.Equal(models.MeasurementTypeDO, measurements[0].Name)

	c.Equal(float64(5.849999904632568), measurements[1].Value)
	c.Equal(models.MeasurementTypePH, measurements[1].Name)

	c.Equal(float64(0), measurements[2].Value)
	c.Equal(models.MeasurementTypeTDS, measurements[2].Name)
}
