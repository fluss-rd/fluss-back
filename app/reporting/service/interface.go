package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/reporting/models"
)

type Service interface {
	GetDetailsReportByModule(ctx context.Context, moduleID string, options models.SearchOptions) (models.Report, error)
}
