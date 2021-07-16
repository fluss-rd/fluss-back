package service

import (
	"context"
	"fmt"
	"log"

	"github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	riverGrpcClient "github.com/flussrd/fluss-back/app/shared/grpc-clients/river-management"
	calculator "github.com/flussrd/fluss-back/app/shared/wqi-calculator"
	"github.com/flussrd/fluss-back/app/telemetry/models"
	repository "github.com/flussrd/fluss-back/app/telemetry/repositories/measurements"
)

var (
	wqiCalculator, _ = calculator.NewCalculator(calculator.IndexTypeWAI)
)

type service struct {
	riverClient *riverGrpcClient.Client
	repo        repository.Repository
}

func New(riverClient *riverGrpcClient.Client, repo repository.Repository) Service {
	return service{
		riverClient: riverClient,
		repo:        repo,
	}
}

func (s service) SaveMeasurement(ctx context.Context, message models.Message) error {
	// TODO: validate message before calling these functions
	client := s.riverClient.GetServiceClient()
	module, err := client.GetModuleByPhonenumber(ctx, &grpchandler.GetModuleRequest{PhoneNumber: message.PhoneNumber})
	if err != nil {
		fmt.Println("getting_module_by_phone_number_failed: ", err.Error())
		return err
	}

	// TODO: move this to a new function
	if module.Currentstate == "inactive" {
		_, err := client.UpdateModuleStatus(ctx, &grpchandler.UpdateModuleRequest{ModuleID: message.ModuleID, Status: "active"})
		if err != nil {
			log.Println("failed to update the module state")
		}
	}

	// TODO: add broken or not field on message

	// Ignore the message
	if module.Currentstate == "deactivated" {
		log.Println("Ignoring message since the module is deactivated")
		return nil
	}

	wqi := wqiCalculator.GetWQI(toCalculatorParams(message.Measurements))

	message.Measurements = append(message.Measurements, models.Measurement{Name: "wqi", Value: wqi})

	message.Measurements = append(message.Measurements, models.Measurement{Name: "lat", Value: module.Location.Latitude})
	message.Measurements = append(message.Measurements, models.Measurement{Name: "lng", Value: module.Location.Longitude})

	message.ModuleID = module.ModuleID
	message.RiverID = module.RiverID

	// TODO: handle the repo errors
	return s.repo.SaveMeasurement(ctx, message)
}

func toCalculatorParams(params []models.Measurement) []calculator.Parameter {
	output := make([]calculator.Parameter, len(params))

	for i := range params {
		output[i].Name = calculator.ParameterType(params[i].Name)
		output[i].Value = params[i].Value
	}

	return output
}
