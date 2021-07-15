package models

import "time"

var validAggregationWindows = map[string]bool{
	"1h": true,
	"1d": true,
	"1m": true,
	"1y": true,
}

type SearchOptions struct {
	RiverID           string    `json:"riverID"`
	Cardinality       string    `json:"cardinality"`
	AggregationWindow string    `json:"aggregationWindow"`
	Start             time.Time `json:"start"`
	End               time.Time `json:"end"`
}

func IsValidAggregationType(aggregationType string) bool {
	return validAggregationWindows[aggregationType]
}
