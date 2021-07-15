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
	permissibleValues = map[ParameterType]float64{
		ParameterTypePH:  7,
		ParameterTypeTDS: 500,
		ParameterTypeTDY: 10, // arbitrary, from Ivan's scale
		ParameterTypeDO:  5,
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
