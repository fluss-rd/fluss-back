package models

import "time"

// Module represents a device on the river, the entity that is going to take the measurements
type Module struct {
	ModuleID     string    `json:"moduleID" bson:"_id"`
	PhoneNumber  string    `json:"phoneNumber" bson:"phoneNumber"`
	RiverID      string    `json:"riverID" bson:"riverID"`
	UserID       string    `json:"userID" bson:"userID"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time `json:"updateDate" bson:"updateDate"`
}
