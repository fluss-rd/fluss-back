package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/accounts/models"
)

var (
	// ErrDuplicateFields duplicate fields
	ErrDuplicateFields = errors.New("duplicate fields")
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// Repository represents the methods of data persistance
type Repository interface {
	GetRole(ctx context.Context, roleName string) (models.Role, error)
	GetRoles(ctx context.Context) ([]models.Role, error)
	CreateRole(ctx context.Context, role models.Role) error
	GetUserRole(ctx context.Context, userID string) ([]models.Role, error)
}
