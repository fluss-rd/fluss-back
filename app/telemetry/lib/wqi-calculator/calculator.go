package calculator

import (
	"errors"

	"github.com/flussrd/fluss-back/app/telemetry/models"
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
	permissibleValues = map[models.MeasurementType]float64{
		models.MeasurementTypePH:  7,
		models.MeasurementTypeTDS: 500,
		models.MeasurementTypeTDY: 10, // arbitrary, from Ivan's scale
		models.MeasurementTypeDO:  5,
	}

	waiUsedValues = map[models.MeasurementType]bool{
		models.MeasurementTypePH:  true,
		models.MeasurementTypeTDS: true,
		models.MeasurementTypeTDY: true, // arbitrary, from Ivan's scale
		models.MeasurementTypeDO:  true,
	}
)

type Calculator interface {
	GetWQI(parameters []models.Measurement) float64
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

func (calculator waiCalculator) GetWQI(parameters []models.Measurement) float64 {
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

func (calculator waiCalculator) shouldUseParam(paramName models.MeasurementType) bool {
	return waiUsedValues[paramName]
}
