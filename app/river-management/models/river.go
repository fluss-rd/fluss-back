package models

type River struct {
	RiverID  string   `json:"riverID" bson:"_id"` // This is "coupling" bewteen the models and mongo
	Name     string   `json:"name" bson:"name"`
	Location Location `json:"location" bson:"location"`
}

type Location struct {
	Lat float64 `json:"latitude" bson:"latitude"`
	Lng float64 `json:"longitude" bson:"latitude"`
}
