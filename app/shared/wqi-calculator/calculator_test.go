package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetWQIStandardValues(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
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
	c.Equal(100.0, wqi)
}

func TestGetWQILowerPh(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 4.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 5,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(100.0, wqi)
}

func TestGetWQIHigherPh(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 12.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 15,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(100.0, wqi)
}

func TestGetWQILowerDO(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 10,
		},
		{
			Name:  "do",
			Value: 1,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(100.0, wqi)
}

func TestGetWQIHigherDO(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 10,
		},
		{
			Name:  "do",
			Value: 10,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(wqi, 100.0)
}

func TestGetWQILowerTDS(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 100,
		},
		{
			Name:  "ty",
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
	c.Less(wqi, 100.0)
}

func TestGetWQIHigherTDS(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 700,
		},
		{
			Name:  "ty",
			Value: 20,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(100.0, wqi)
}

func TestGetWQILowerTY(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 2,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(wqi, 100.0)
}

func TestGetWQIHigherTY(t *testing.T) {
	c := require.New(t)

	measurements := []Parameter{
		{
			Name:  "ph",
			Value: 7.5,
		},
		{
			Name:  "tds",
			Value: 500,
		},
		{
			Name:  "ty",
			Value: 100,
		},
		{
			Name:  "do",
			Value: 5,
		},
	}

	calculator, err := NewCalculator(IndexTypeWAI)
	c.Nil(err)

	wqi := calculator.GetWQI(measurements)
	c.Less(100.0, wqi)
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
