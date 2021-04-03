package models

import "time"

// User user represents a user on the platform
type User struct {
	UserID       string    `bson:"_id"`
	PhoneNumber  string    `bson:"phoneNumber"`
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	Password     string    `bson:"password"`
	RoleName     string    `bson:"roleName"`
	CreationDate time.Time `bson:"creationDate"`
	UpdateDate   time.Time `bson:"updateDate"`
}
