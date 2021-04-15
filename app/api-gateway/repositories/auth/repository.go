package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/accounts/models"
)

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// Repository authentication and authorization repository
type Repository interface {
	GetRole(ctx context.Context, roleName string) (models.Role, error) // TODO: move that role to a shared package or create one here
}
