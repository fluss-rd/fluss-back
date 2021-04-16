package models

import "time"

// River represents a physical river which is going to be measured
type River struct {
	RiverID      string    `json:"riverID" bson:"_id"` // This is "coupling" bewteen the models and mongo
	Name         string    `json:"name" bson:"name"`
	Location     Location  `json:"location" bson:"location"`
	UserID       string    `json:"userID" bson:"userID"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time `json:"updateDate" bson:"updateDate"`
}

// Location location of the entity
type Location struct {
	Lat float64 `json:"latitude" bson:"latitude"`
	Lng float64 `json:"longitude" bson:"longitude"`
}
