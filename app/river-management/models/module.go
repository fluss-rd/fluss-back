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
	// ModuleStateInactive the default state of a module when its created but hasnt started sending data
	ModuleStateInactive = "inactive"
	// ModuleStateDeactivated when the module was deactivated and we will not process further messages
	ModuleStateDeactivated = "deactivated"
	// ModuleStateDeleted the module was deleted
	ModuleStateDeleted = "deleted"
	// ModuleStateBroken the module has been presenting technical issues
	ModuleStateBroken = "broken"
)

// Module represents a device on the river, the entity that is going to take the measurements
type Module struct {
	ModuleID     string      `json:"moduleID" bson:"_id"`
	PhoneNumber  string      `json:"phoneNumber" bson:"phoneNumber"`
	Alias        string      `json:"alias" bson:"alias"`
	// TODO: RiverID should be WaterBodyID
	RiverID      string      `json:"riverID" bson:"riverID"`
	// TODO: RiverName should be WaterBodyName
	RiverName    string      `json:"riverName" bson:"riverName"`
	UserID       string      `json:"userID" bson:"userID"`
	CreationDate time.Time   `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time   `json:"updateDate" bson:"updateDate"`
	CurrentState ModuleState `json:"state" bson:"state"`
	Serial       string      `json:"serial" bson:"serial"`
	Location     Point       `json:"location" bson:"location"`
}

type ModuleUpdateOptions struct {
	Status   ModuleState
	ModuleID string
}
