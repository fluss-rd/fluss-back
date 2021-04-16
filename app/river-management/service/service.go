package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/accounts/shared/utils"
	"github.com/flussrd/fluss-back/app/river-management/models"
	modulesRepository "github.com/flussrd/fluss-back/app/river-management/repositories/modules"
	riversRepository "github.com/flussrd/fluss-back/app/river-management/repositories/rivers"
)

var (
	// ErrMissingName missing name
	ErrMissingName = httputils.NewBadRequestError("missing name")
	// ErrMissingUserID missing user id
	ErrMissingUserID = httputils.NewBadRequestError("missing user id")
	// ErrMissingLatitude missing latitude
	ErrMissingLatitude = httputils.NewBadRequestError("missing latitude")
	// ErrMissingLongitude missing longitude
	ErrMissingLongitude = httputils.NewBadRequestError("missing longitude")
	// ErrMissingPhoneNumber missing phone number
	ErrMissingPhoneNumber = httputils.NewBadRequestError("missing phone number")
	// ErrMissingRiverID missing river id
	ErrMissingRiverID = httputils.NewBadRequestError("missing river id")
	// ErrMissingModuleID missing module id
	ErrMissingModuleID = httputils.NewBadRequestError("missing module id")

	// ErrGeneratingIDFailed generating id failed
	ErrGeneratingIDFailed = errors.New("generating id failed")
	// ErrSavingRiverFailed saving river failed
	ErrSavingRiverFailed = errors.New("saving river failed")
)

type service struct {
	riversRepo  riversRepository.Repository
	modulesRepo modulesRepository.Repository
}

// New returns a new service
func New(riversRepo riversRepository.Repository, modulesRepo modulesRepository.Repository) Service {
	return service{
		riversRepo:  riversRepo,
		modulesRepo: modulesRepo,
	}
}

func (s service) CreateRiver(ctx context.Context, river models.River) (models.River, error) {
	err := validateCreateRiverFields(river)
	if err != nil {
		return models.River{}, nil
	}

	id, err := utils.GenerateID("RVR")
	if err != nil {
		return models.River{}, fmt.Errorf("%w: %s", ErrGeneratingIDFailed, err.Error())
	}

	// TODO: validate if the user ID exists consuming the accounts service. we should create a client library

	river.RiverID = id

	err = s.riversRepo.SaveRiver(ctx, river)
	if err != nil {
		return models.River{}, fmt.Errorf("%w: %s", ErrSavingRiverFailed, err.Error())
	}

	return river, nil
}

func validateCreateRiverFields(river models.River) error {
	if river.Name == "" {
		return ErrMissingName
	}

	if river.UserID == "" {
		return ErrMissingUserID
	}

	if river.Location.Lat == 0 {
		return ErrMissingLatitude
	}

	if river.Location.Lng == 0 {
		return ErrMissingLongitude
	}

	return nil
}

func (s service) GetRiversN(ctx context.Context) ([]models.River, error) {
	rivers, err := s.riversRepo.GetAllRiversNotPaginated(ctx)
	if errors.Is(err, riversRepository.ErrNotFound) {
		return nil, httputils.NewNotFoundError("river")
	}

	if err != nil {
		// TODO: wrap error
		return nil, err
	}

	return rivers, nil
}

func (s service) CreateModule(ctx context.Context, module models.Module) (models.Module, error) {
	err := validateCreateModuleFields(module)
	if err != nil {
		return models.Module{}, err
	}

	// Modules become active as soon we receive the first data coming from the. Until then, the module is inactive
	module.CurrentState = models.ModuleStateInactive

	id, err := utils.GenerateID("MDL")
	if err != nil {
		return models.Module{}, ErrGeneratingIDFailed
	}

	module.ModuleID = id

	return module, nil
}

func validateCreateModuleFields(module models.Module) error {
	if module.PhoneNumber == "" {
		return ErrMissingPhoneNumber
	}

	if module.RiverID == "" {
		return ErrMissingRiverID
	}

	if module.UserID == "" {
		return ErrMissingUserID
	}

	return nil
}

func (s service) GetModule(ctx context.Context, moduleID string) (models.Module, error) {
	if moduleID == "" {
		return models.Module{}, ErrMissingModuleID
	}

	module, err := s.GetModule(ctx, moduleID)
	if errors.Is(err, modulesRepository.ErrNotFound) {
		return models.Module{}, httputils.NewNotFoundError("module")
	}

	return module, nil
}

func (s service) GetModulesN(ctx context.Context) ([]models.Module, error) {
	modules, err := s.modulesRepo.GetAllModulesWithOutPagination(ctx)
	if errors.Is(err, modulesRepository.ErrNotFound) {
		return []models.Module{}, nil
	}

	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (s service) GetModulesByRiverN(ctx context.Context, riverID string) ([]models.Module, error) {
	return nil, nil
}
