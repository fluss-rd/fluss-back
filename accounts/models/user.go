package models

import "time"

// User user represents a user on the platform
type User struct {
	UserID       string
	PhoneNumber  string
	Name         string
	Email        string
	Password     string
	RoleName     string
	CreationDate time.Time
	UpdateDate   time.Time
}
