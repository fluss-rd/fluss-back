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

	waiUsedValues = map[ParameterType]bool{
		ParameterTypePH:  true,
		ParameterTypeTDS: true,
		ParameterTypeTDY: true, // arbitrary, from Ivan's scale
		ParameterTypeDO:  true,
	}
)

type Calculator interface {
	GetWQI(parameters []Parameter) float64
}

// Weighted Arithmetic Index
// Found on "Analytical Studies on Water Quality Index of River Landzu"
type waiCalculator struct {
}

func NewCalculator(indexType IndexType) (Calculator, error) {
	switch indexType {
	case IndexTypeWAI:
		return waiCalculator{}, nil
	}

	return nil, ErrInvalidIndexType
}

func (calculator waiCalculator) GetWQI(parameters []Parameter) float64 {
	wqSum := 0.0
	wSum := 0.0

	// quality rating scale. generation of the parameter sub-indices: parameter concentrations are converted to unit less sub-indices
	for _, param := range parameters {
		if !calculator.shouldUseParam(param.Name) {
			continue
		}

		permissibleValue := permissibleValues[param.Name]
		w := (1 / permissibleValue)

		q := (param.Value / permissibleValue) * 100

		// agregation
		wSum += w
		wqSum += w * q
	}

	return wqSum / wSum // overall
}

func (calculator waiCalculator) shouldUseParam(paramName ParameterType) bool {
	return waiUsedValues[paramName]
}
