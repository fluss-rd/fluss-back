package repository

import (
	"context"

	"github.com/flussrd/fluss-back/app/reporting/models"
)

type Parameter struct {
	Name  string
	Value float64
}

type Repository interface {
	GetDataByModule(ctx context.Context, moduleID string, searchOptions models.SearchOptions) (models.Report, error)
}
