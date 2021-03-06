package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/accounts/models"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
)

// UserUseCase defines the usecases functions for the user
type UserUseCase interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	AddRoleToUser(ctx context.Context, roleName string, userID string) error
	UpdateUser(ctx context.Context, request httputils.PatchRequest, userID string) (models.User, error)
	GetUser(ctx context.Context, userID string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}

// RolesUseCase defines the usescases functions for the user
type RolesUseCase interface {
	CreateRole(ctx context.Context, role models.Role) error
	GetRoles(ctx context.Context) ([]models.Role, error)
	UpdateRole(ctx context.Context, role models.Role) error
}

// AuthUseCase defines the usecases function for the auth operations
type AuthUseCase interface {
	Login(ctx context.Context, email, password string) (LoginResponse, error)
}

// Service defines the service layer methods
type Service interface {
	RolesUseCase
	UserUseCase
	AuthUseCase
}
