package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/reporting/models"
	repository "github.com/flussrd/fluss-back/app/reporting/repositories/reports"
	calculator "github.com/flussrd/fluss-back/app/shared/wqi-calculator"
)

var (
	// TODO: error handling?
	indexCalculator, _ = calculator.NewCalculator(calculator.IndexTypeWAI)
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
	report, err := s.repo.GetDataByModule(ctx, moduleID, options)
	if err != nil {
		return models.Report{}, err
	}

	reports := addWQIToReports([]models.Report{report})

	return reports[0], nil
}

func (s service) GetAllModulesSummary(ctx context.Context, options models.SearchOptions) ([]models.Report, error) {
	reports, err := s.repo.GetAllModulesSummary(ctx, options)
	if err != nil {
		return nil, err
	}

	reports = addWQIToReports(reports)

	return reports, nil
}

func toCalculatorParams(params []models.Parameter) []calculator.Parameter {
	output := make([]calculator.Parameter, len(params))

	for i := range params {
		output[i].Name = params[i].Name
		output[i].Value = params[i].Value
	}

	return output
}

func addWQIToReports(reports []models.Report) []models.Report {
	for index, report := range reports {
		for j, data := range report.Data {
			wqi := indexCalculator.GetWQI(toCalculatorParams(data.Parameters))
			reports[index].Data[j].WQI = wqi
			reports[index].Data[j].WQIClassification = indexCalculator.GetWQIClassification(wqi)
		}
	}

	return reports
}
