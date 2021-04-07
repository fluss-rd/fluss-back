package service

import (
	"context"

	"github.com/flussrd/fluss-back/accounts/models"
)

// UserUseCase defines the usecases functions for the user
type UserUseCase interface {
	CreateUser(ctx context.Context, user models.User) error
	AddRoleToUser(ctx context.Context, roleName string, userID string) error
}

// RolesUseCase defines the usescases functions for the user
type RolesUseCase interface {
	CreateRole(ctx context.Context, role models.Role) error
	UpdateRole(ctx context.Context, role models.Role) error
}

// Service defines the service layer methods
type Service interface {
	RolesUseCase
	UserUseCase
}
