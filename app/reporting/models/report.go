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
	LastDate   time.Time   `json:"lastDate"`
	Parameters []Parameter `json:"parameters"`
	Location   Location    `json:"location"`
}

type Parameter struct {
	Name  string    `json:"name"`
	Value float64   `json:"value"`
	Date  time.Time `json:"date"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
