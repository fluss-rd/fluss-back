package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/river-management/models"
)

// RiversUsecases defines the use cases functions related to rivers
type RiversUsecases interface {
	CreateRiver(ctx context.Context, river models.River) (models.River, error)
	GetRiversN(ctx context.Context) ([]models.River, error) //  no pagination
	GetRiver(ctx context.Context, riverID string) (models.River, error)
}

// ModulesUsecases defines the use cases functions related to modules
type ModulesUsecases interface {
	CreateModule(ctx context.Context, module models.Module) (models.Module, error)
	GetModule(ctx context.Context, moduleID string) (models.Module, error)
	GetModuleByPhoneNumber(ctx context.Context, phoneNumber string) (models.Module, error)
	GetModulesN(ctx context.Context) ([]models.Module, error)                        // no pagination
	GetModulesByRiverN(ctx context.Context, riverID string) ([]models.Module, error) // no pagination
}

// Service defines the method for the service/usecase layer
type Service interface {
	RiversUsecases
	ModulesUsecases
}
