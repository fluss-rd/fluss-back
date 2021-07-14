package models

import "time"

type Report struct {
	ModuleID    string    `json:"moduleID"`
	RiverID     string    `json:"riverID"`
	WQI         float64   `json:"wqi"`
	PH          float64   `json:"ph"`
	TDS         float64   `json:"tds"`
	DO          float64   `json:"do"`
	TMP         float64   `json:"tmp"`
	Location    float64   `json:"location"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
