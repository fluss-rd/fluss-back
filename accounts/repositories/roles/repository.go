package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/accounts/models"
)

var (
	ErrDuplicateFields = errors.New("duplicate fields")
)

type Repository interface {
	GetRole(ctx context.Context, roleName string) (models.Role, error)
	CreateRole(ctx context.Context, role models.Role) models.Role
	GetUserRole(ctx context.Context, userID string) ([]models.Role, error)
}
