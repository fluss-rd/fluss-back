package handlers

import (
	"context"
	"fmt"

	grpcService "github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	"github.com/flussrd/fluss-back/app/river-management/service"
)

type Handler interface {
	GetModuleByPhonenumber(ctx context.Context, request *grpcService.GetModuleRequest) (*grpcService.Module, error)
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
	fmt.Println("called")
	module, err := handler.s.GetModuleByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
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
