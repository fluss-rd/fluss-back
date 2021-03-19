package models

type ResourceType string

const (
	ResourceTypeModule ResourceType = "module"
	ResourceTypeUser   ResourceType = "user"
)

type Role struct {
	Name        string       `json:"roleName" bson:"roleName"`
	Permissions []Permission `bson:"permissions"`
}

type Permission struct {
	Resource ResourceType `bson:"resource"`
	Action   string       `bson:"action"`
}
