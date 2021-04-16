package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/river-management/models"
)

type RiversUsecases interface {
	CreateRiver(ctx context.Context, river models.River) error
	GetRiversN(ctx context.Context) ([]models.River, error) //  no pagination
}

type ModulesUsecases interface {
	CreateModule(ctx context.Context, module models.Module) error
	GetModule(ctx context.Context, moduleID string) (models.Module, error)
	GetModulesN(ctx context.Context) ([]models.Module, error)                        // no pagination
	GetModulesByRiverN(ctx context.Context, riverID string) ([]models.Module, error) // no pagination
}

type Service interface {
	RiversUsecases
	ModulesUsecases
}
