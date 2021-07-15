package models

type MeasurementType string

const (
	// TODO: maybe change this to parameter?
	// MeasurementTypePH is a MeasurementType for the pH level of the water.
	MeasurementTypePH MeasurementType = "ph"
	// MeasurementTypeDO is a MeasurementType for the concentration of oxygen that is dissolved in the water (Dissolved Oxygen).
	MeasurementTypeDO  MeasurementType = "do"
	// MeasurementTypeTMP is a MeasurementType for the temperature of the water.
	MeasurementTypeTMP MeasurementType = "tmp"
	// MeasurementTypeTDY is a MeasurementType for the turbidity of the water.
	MeasurementTypeTDY MeasurementType = "ty"
	// MeasurementTypeTDS is a MeasurementType for the total dissolved solid (TDS) of the water.
	MeasurementTypeTDS MeasurementType = "tds"
	// MeasurementTypeLat is a MeasurementType for the latitude of the module location.
	MeasurementTypeLat MeasurementType = "lat"
	// MeasurementTypeLng is a MeasurementType for the longitude of the module location.
	MeasurementTypeLng MeasurementType = "lng"
)

type Measurement struct {
	Name  MeasurementType `json:"name"`
	Value float64         `json:"value"`
}
