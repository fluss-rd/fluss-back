package service

import (
	"context"

	"github.com/flussrd/fluss-back/accounts/models"
	rolesRepository "github.com/flussrd/fluss-back/accounts/repositories/roles"
	usersRepository "github.com/flussrd/fluss-back/accounts/repositories/users"
)

type service struct {
	usersRepo usersRepository.Repository
	rolesRepo rolesRepository.Repository
}

// NewService returns a new service entity to be able to execuse business logic
func NewService(usersRepo usersRepository.Repository, rolesRepo rolesRepository.Repository) Service {
	return service{
		usersRepo: usersRepo,
		rolesRepo: rolesRepo,
	}
}

// CreateUser creates a new user
func (s service) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

// AddRoleToUser adds a role to a user
func (s service) AddRoleToUser(ctx context.Context, roleName string, userID string) error {
	return nil
}

// CreateRole creates a new role
func (s service) CreateRole(ctx context.Context, role models.Role) error {
	return nil
}

// UpdateRole updates a role
func (s service) UpdateRole(ctx context.Context, role models.Role) error {
	return nil
}
