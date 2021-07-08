package models

import "time"

// BodyType represents the body of water type.
type BodyType string

const (
	// BodyTypeRiver river
	BodyTypeRiver BodyType = "river"
	// BodyTypeLagoon lagoon body of water type
	BodyTypeLagoon BodyType = "lagoon"
	// BodyTypeLake lake body of water type
	BodyTypeLake BodyType = "lake"
	// BodyTypeStream stream body of water type
	BodyTypeStream BodyType = "stream"
)

var validBodytTypes = map[BodyType]bool{
	BodyTypeRiver:  true,
	BodyTypeLagoon: true,
	BodyTypeLake:   true,
	BodyTypeStream: true,
}

// River represents a physical river which is going to be measured
type River struct {
	RiverID      string    `json:"riverID" bson:"_id"` // This is "coupling" bewteen the models and mongo
	Name         string    `json:"name" bson:"name"`
	Location     Location  `json:"location" bson:"location"`
	UserID       string    `json:"userID" bson:"userID"`
	Type         BodyType  `json:"type" bson:"type"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time `json:"updateDate" bson:"updateDate"`
}

// Location location of the entity
type Location struct {
	Lat float64 `json:"latitude" bson:"latitude"`
	Lng float64 `json:"longitude" bson:"longitude"`
}

// IsValidBodyType returns whether a given body of water type is valid or not
func IsValidBodyType(bodyType BodyType) bool {
	return validBodytTypes[bodyType]
}
