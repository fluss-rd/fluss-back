package calculator

type ParameterType string

const (
	// TODO: maybe change this to parameter?
	// ParameterTypePH is a ParameterType for the pH level of the water.
	ParameterTypePH ParameterType = "ph"
	// ParameterTypeDO is a ParameterType for the concentration of oxygen that is dissolved in the water (Dissolved Oxygen).
	ParameterTypeDO ParameterType = "do"
	// ParameterTypeTMP is a ParameterType for the temperature of the water.
	ParameterTypeTMP ParameterType = "tmp"
	// ParameterTypeTDY is a ParameterType for the turbidity of the water.
	ParameterTypeTDY ParameterType = "ty"
	// ParameterTypeTDS is a ParameterType for the total dissolved solid (TDS) of the water.
	ParameterTypeTDS ParameterType = "tds"
	// ParameterTypeLat is a ParameterType for the latitude of the module location.
	ParameterTypeLat ParameterType = "lat"
	// ParameterTypeLng is a ParameterType for the longitude of the module location.
	ParameterTypeLng ParameterType = "lng"
)

type Parameter struct {
	Name  ParameterType `json:"name"`
	Value float64       `json:"value"`
}
