package models

type MeasurementType string

const (
	// TODO: maybe change this to parameter?
	MeasurementTypePH MeasurementType = "ph"
	// MeasurementTypeOx oxygen
	MeasurementTypeDO  MeasurementType = "do"
	MeasurementTypeTMP MeasurementType = "tmp"
	// MeasurementTypeTY turbidity
	MeasurementTypeTDY MeasurementType = "ty"
	MeasurementTypeTDS MeasurementType = "tds"
	MeasurementTypeTO  MeasurementType = "do"
	MeasurementTypeLat MeasurementType = "lat"
	MeasurementTypeLng MeasurementType = "lng"
)

type Measurement struct {
	Name  MeasurementType `json:"name"`
	Value float64         `json:"value"`
}
