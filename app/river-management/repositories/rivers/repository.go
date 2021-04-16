package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/river-management/models"
)

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// Repository defines the data-persistance methods for rivers
type Repository interface {
	SaveRiver(ctx context.Context, river models.River) (models.River, error)
	GetRiver(ctx context.Context, riverID string) (models.River, error)
	GetAllRiversNotPaginated(ctx context.Context) ([]models.River, error) // This will be replaced for a paginated one
	GetAllRivers(ctx context.Context) ([]models.River, string, error)
}
