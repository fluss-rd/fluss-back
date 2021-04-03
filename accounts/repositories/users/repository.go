package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/accounts/models"
)

var (
	// ErrDuplicateFields duplicate fields
	ErrDuplicateFields = errors.New("duplicate fields")
)

// Repository represents the methods of data persistance
type Repository interface {
	GetUser(ctx context.Context, userID string) (models.User, error)
	SaveUser(ctx context.Context, user models.User) (models.User, error)
	AddRoleToUser(ctx context.Context, userID string, role models.Role) error
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
}
