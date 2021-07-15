package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetWQI(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
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

func TestGetWQIClassification(t *testing.T) {
	c := require.New(t)

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	c.Equal("excellent", calculator.GetWQIClassification(40))
	c.Equal("good water", calculator.GetWQIClassification(70))
	c.Equal("poor water", calculator.GetWQIClassification(150))
	c.Equal("very poor water", calculator.GetWQIClassification(250))
	c.Equal("unsuitable", calculator.GetWQIClassification(450))
}
