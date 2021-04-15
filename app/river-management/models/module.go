package models

// Module
type Module struct {
	ModuleID string `json:"moduleID" bson:"_id"`
	RiverID  string `json:"riverID" bson:"riverID"`
}
