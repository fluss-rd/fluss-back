package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/reporting/models"
)

type Service interface {
	GetAllModulesSummary(ctx context.Context, options models.SearchOptions) ([]models.Report, error)
	GetDetailsReportByModule(ctx context.Context, moduleID string, options models.SearchOptions) (models.Report, error)
	GetRiverSummary(ctx context.Context, riverID string) (models.Report, error)
}
