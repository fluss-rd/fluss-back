package calculator

import (
	"math"
)

var (
	waiUsedValues = map[ParameterType]bool{
		ParameterTypePH:  true,
		ParameterTypeTDS: true,
		ParameterTypeTDY: true,
		ParameterTypeDO:  true,
	}
)

// Weighted Arithmetic Index
// Found on "Analytical Studies on Water Quality Index of River Landzu"
type waiCalculator struct {
}

func (calculator waiCalculator) GetWQI(parameters []Parameter) float64 {
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

func (calculator waiCalculator) shouldUseParam(paramName ParameterType) bool {
	return waiUsedValues[paramName]
}

func (calculator waiCalculator) GetWQIClassification(wqi float64) string {
	if wqi < 50 {
		return "excellent"
	} else if wqi < 100 {
		return "good water"
	} else if wqi < 200 {
		return "poor water"
	} else if wqi < 300 {
		return "very poor water"
	}

	return "unsuitable"
}
