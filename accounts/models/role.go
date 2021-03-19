package models

type ResourceType string

const (
	ResourceTypeModule ResourceType = "module"
	ResourceTypeUser   ResourceType = "user"
)

type Role struct {
	Name        string
	Permissions []Permission
}

type Permission struct {
	Resource ResourceType
	Action   string
}
