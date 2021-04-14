package models

import "time"

// User user represents a user on the platform
type User struct {
	UserID       string    `json:"userID" bson:"_id"`
	PhoneNumber  string    `json:"phoneNumber" bson:"phoneNumber"`
	Name         string    `json:"name" bson:"name"`
	Email        string    `json:"email" bson:"email"`
	Password     string    `json:"password" bson:"password"`
	RoleName     string    `json:"roleName" bson:"roleName"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	UpdateDate   time.Time `json:"updateDate" bson:"updateDate"`
}
