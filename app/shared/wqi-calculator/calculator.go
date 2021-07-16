package calculator

import (
	"errors"
)

type IndexType string

const (
	// IndexTypeWAI wai
	IndexTypeWAI IndexType = "wai"
)

var (
	ErrInvalidIndexType = errors.New("invalid index type")
)

var (
	// From "Analytical Studies on Water Quality Index of River Landzu", table 2
	standardValues = map[ParameterType]float64{
		ParameterTypePH:  7.5,
		ParameterTypeTDS: 500, // measured in mg/litre
		ParameterTypeTDY: 10, // measured in NTUs
		ParameterTypeDO:  5, // measured in mg/litre
	}

	idealValues = map[ParameterType]float64{
		ParameterTypePH:  7,
		ParameterTypeTDS: 0,
		ParameterTypeTDY: 0,
		ParameterTypeDO:  20, // 20 is the max the sensor can read
	}
)

type Calculator interface {
	GetWQI(parameters []Parameter) float64
	GetWQIClassification(wqi float64) string
}

func NewCalculator(indexType IndexType) (Calculator, error) {
	switch indexType {
	case IndexTypeWAI:
		return waiCalculator{}, nil
	}

	return nil, ErrInvalidIndexType
}
