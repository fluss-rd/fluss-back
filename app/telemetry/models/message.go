package models

import "time"

type MessageType string

const (
	MessageTypeMeasurement MessageType = "measurement"
)

type Message struct {
	ID           string        `json:"id"`
	Date         time.Time     `json:"date"`
	ModuleID     string        `json:"moduleID"`
	PhoneNumber  string        `json:"phoneNumber"`
	MessageType  MessageType   `json:"messageType"`
	Measurements []Measurement `json:"measurements,omitempty"`
}
