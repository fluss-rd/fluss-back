package models

import "time"

type User struct {
	UserID       string
	PhoneNumber  string
	Name         string
	Email        string
	Password     string
	CreationDate time.Time
	UpdateDate   time.Time
}
