package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/river-management/models"
)

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
	// ErrDuplicateFields duplicate fields
	ErrDuplicateFields = errors.New("duplicate fields")
)

// Repository defines the data-persistance methods related to modules
type Repository interface {
	GetModule(ctx context.Context, moduleID string) (models.Module, error)
	GetAllModulesWithOutPagination(ctx context.Context) ([]models.Module, error)
	GetAllModules(ctx context.Context) ([]models.Module, string, error)
	GetModulesByRiverWithoutPagination(ctx context.Context, riverID string) ([]models.Module, error)
	GetModulesByRiver(ctx context.Context) ([]models.Module, string, error)
	SaveModule(ctx context.Context, module models.Module) (models.Module, error)
}
