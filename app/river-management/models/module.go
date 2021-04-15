package models

import "time"

// ModuleState defines the module state
// Is one of:
// - active
// - inactive
// - deleted
// - broken
type ModuleState string

const (
	// ModuleStateActive the module is fully working and reporting data
	ModuleStateActive = "active"
	// ModuleStateInactive the module was deactiviated and data sent from it will not be stored
	ModuleStateInactive = "inactive"
	// ModuleStateDeleted the module was deleted
	ModuleStateDeleted = "deleted"
	// ModuleStateBroken the module has been presenting technical issues
	ModuleStateBroken = "broken"
)

// Module represents a device on the river, the entity that is going to take the measurements
type Module struct {
	ModuleID     string      `json:"moduleID" bson:"_id"`
	PhoneNumber  string      `json:"phoneNumber" bson:"phoneNumber"`
	RiverID      string      `json:"riverID" bson:"riverID"`
	UserID       string      `json:"userID" bson:"userID"`
	CreationDate time.Time   `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time   `json:"updateDate" bson:"updateDate"`
	CurrentState ModuleState `json:"state" bson:"state"`
	Serial       string      `json:"serial" bson:"serial"`
}
