package models

type SearchOptions struct {
	Cardinality       string `json:"cardinality"`
	AggregationWindow string `json:"aggregationWindow"`
}
