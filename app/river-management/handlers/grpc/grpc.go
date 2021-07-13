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
		return nil, err
	}

	return &grpcService.Module{
		ModuleID:    module.ModuleID,
		PhoneNumber: module.PhoneNumber,
	}, nil
}
