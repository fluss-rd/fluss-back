package models

import (
	"time"

	calculator "github.com/flussrd/fluss-back/app/shared/wqi-calculator"
)

type Report struct {
	ModuleID    string    `json:"moduleID"`
	RiverID     string    `json:"riverID"`
	Data        []Data    `json:"data"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Data struct {
	WQI               float64     `json:"wqi"`
	WQIClassification string      `json:"wqiClassification"`
	LastDate          time.Time   `json:"lastDate"`
	Parameters        []Parameter `json:"parameters"`
	Location          Location    `json:"location"`
}

type Parameter struct {
	calculator.Parameter
	Date time.Time `json:"date"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
