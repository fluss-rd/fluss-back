package models

import "time"

type Report struct {
	ModuleID    string    `json:"moduleID"`
	RiverID     string    `json:"riverID"`
	Data        []Data    `json:"data"`
	Location    float64   `json:"location"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Data struct {
	Date       time.Time   `json:"date"`
	Parameters []Parameter `json:"parameters"`
	Location   Location    `json:"location"`
}

type Parameter struct {
	Name  string
	Value float64
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
