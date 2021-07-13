package models

type MeasurementType string

const (
	MeasurementTypePH MeasurementType = "ph"
	// MeasurementTypeOx oxygen
	MeasurementTypeOx  MeasurementType = "ox"
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
