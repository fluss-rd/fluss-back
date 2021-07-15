package calculator

var (
	waiUsedValues = map[ParameterType]bool{
		ParameterTypePH:  true,
		ParameterTypeTDS: true,
		ParameterTypeTDY: true, // arbitrary, from Ivan's scale
		ParameterTypeDO:  true,
	}
)

// Weighted Arithmetic Index
// Found on "Analytical Studies on Water Quality Index of River Landzu"
type waiCalculator struct {
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
