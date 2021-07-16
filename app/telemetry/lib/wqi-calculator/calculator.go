package calculator

import (
	"errors"
	"math"

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
	idealValues = map[models.MeasurementType]float64{
		models.MeasurementTypePH:  7,
		models.MeasurementTypeTDS: 0,
		models.MeasurementTypeTDY: 0,
		models.MeasurementTypeDO:  20, // 20 is the max the sensor can read
	}

	// From "Analytical Studies on Water Quality Index of River Landzu", table 2
	standardValues = map[models.MeasurementType]float64{
		models.MeasurementTypePH:  7.5,
		models.MeasurementTypeTDS: 500, // measured in mg/litre
		models.MeasurementTypeTDY: 10, // measured in NTUs
		models.MeasurementTypeDO:  5, // measured in mg/litre
	}

	waiUsedValues = map[models.MeasurementType]bool{
		models.MeasurementTypePH:  true,
		models.MeasurementTypeTDS: true,
		models.MeasurementTypeTDY: true,
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
	wSum := 0.0
	OverallWQI := 0.0
	k := 0.0

	for measurementType, value := range standardValues {
		if !calculator.shouldUseParam(measurementType) {
			continue
		}

		w := (1.0 / value)
		wSum += w
	}

	k = 1.0 / wSum

	// quality rating scale. generation of the parameter sub-indices: parameter concentrations are converted to unit less sub-indices
	for _, param := range parameters {
		if !calculator.shouldUseParam(param.Name) {
			continue
		}

		qi := ( math.Abs(param.Value - idealValues[param.Name]) ) / ( math.Abs(standardValues[param.Name] - idealValues[param.Name]) ) * 100.0
		wi := k / standardValues[param.Name]

		OverallWQI += (wi * qi)
	}

	return OverallWQI
}

func (calculator waiCalculator) shouldUseParam(paramName models.MeasurementType) bool {
	return waiUsedValues[paramName]
}
