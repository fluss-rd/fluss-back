package calculator

import (
	"testing"

	"github.com/flussrd/fluss-back/app/telemetry/models"
	"github.com/stretchr/testify/require"
)

func TestGetWQI(t *testing.T) {
	c := require.New(t)

	measurements := []models.Measurement{
		{
			Name:  "ph",
			Value: 7,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "tdy",
			Value: 10,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Equal(99.99999999999999, wqi)
}
