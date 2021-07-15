package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/reporting/models"
	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return service{
		repo: repo,
	}
}

func (s service) GetDetailsReportByModule(ctx context.Context, moduleID string, options models.SearchOptions) (models.Report, error) {
	return s.repo.GetDataByModule(ctx, moduleID, options)
}
