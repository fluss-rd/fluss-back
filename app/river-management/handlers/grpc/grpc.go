package handlers

import (
	"context"
	"fmt"
	"log"

	grpcService "github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	"github.com/flussrd/fluss-back/app/river-management/models"
	"github.com/flussrd/fluss-back/app/river-management/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler interface {
	GetModuleByPhonenumber(ctx context.Context, request *grpcService.GetModuleRequest) (*grpcService.Module, error)
	UpdateModuleStatus(ctx context.Context, request *grpcService.UpdateModuleRequest) (*emptypb.Empty, error)
}

type grpcHandler struct {
	s service.Service
}

func NewHandler(s service.Service) Handler {
	return grpcHandler{
		s: s,
	}
}

func (handler grpcHandler) GetModuleByPhonenumber(ctx context.Context, request *grpcService.GetModuleRequest) (*grpcService.Module, error) {
	module, err := handler.s.GetModuleByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		log.Println("failed to get module by phone number: ", err.Error())
		return nil, fmt.Errorf("%w with phone number: %s", err, request.PhoneNumber)
	}

	return &grpcService.Module{
		ModuleID:    module.ModuleID,
		PhoneNumber: module.PhoneNumber,
		RiverID:     module.RiverID,
		Location: &grpcService.Point{
			Latitude:  module.Location.Lat,
			Longitude: module.Location.Lng,
		},
	}, nil
}

func (handler grpcHandler) UpdateModuleStatus(ctx context.Context, request *grpcService.UpdateModuleRequest) (*emptypb.Empty, error) {
	_, err := handler.s.UpdateModule(ctx, request.ModuleID, models.ModuleUpdateOptions{State: models.ModuleState(request.Status)})
	if err != nil {
		log.Println("failed to update module status", err.Error())
	}

	return &emptypb.Empty{}, err
}
