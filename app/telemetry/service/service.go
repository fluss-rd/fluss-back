package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"

	"github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	riverGrpcClient "github.com/flussrd/fluss-back/app/shared/grpc-clients/river-management"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
	calculator "github.com/flussrd/fluss-back/app/shared/wqi-calculator"
	"github.com/flussrd/fluss-back/app/telemetry/models"
	repository "github.com/flussrd/fluss-back/app/telemetry/repositories/measurements"
)

var (
	wqiCalculator, _ = calculator.NewCalculator(calculator.IndexTypeWAI)

	parametersRegexsMatch = map[models.MeasurementType]*regexp.Regexp{
		models.MeasurementTypeTMP: regexp.MustCompile(`TEMP\*([+-]?([0-9]*[.])?[0-9]+)`),
		models.MeasurementTypePH:  regexp.MustCompile(`pH\?([+-]?([0-9]*[.])?[0-9]+)`),
		models.MeasurementTypeTDS: regexp.MustCompile(`TDS\+([+-]?([0-9]*[.])?[0-9]+)`),
		models.MeasurementTypeTDY: regexp.MustCompile(`TURB!([+-]?([0-9]*[.])?[0-9]+)`),
		models.MeasurementTypeDO:  regexp.MustCompile(`\(D\.O%([+-]?([0-9]*[.])?[0-9]+)`),
	}
)

type service struct {
	riverClient  *riverGrpcClient.Client
	repo         repository.Repository
	rabbitClient rabbit.RabbitClient
}

func New(riverClient *riverGrpcClient.Client, repo repository.Repository, rabbitClient rabbit.RabbitClient) Service {
	return service{
		riverClient:  riverClient,
		repo:         repo,
		rabbitClient: rabbitClient,
	}
}

// TODO: if we decide this function will handle the validation(headers/host/etc) this should also receve the request
func (s service) HandleHTTPMessage(ctx context.Context, source string, body string) {
	fmt.Println("source: " + source)
	switch source {
	case "twilio":
		s.handleTwilioMessage(ctx, body)
		return
	default:
		log.Println("not_supported_source_message_arrived")
	}
}

func (s service) handleTwilioMessage(ctx context.Context, body string) {
	params, err := url.ParseQuery(body)
	if err != nil {
		log.Println("parsing_message_body_failed: %w", err)
		return
	}

	measurements, err := s.getMeasurementsFromMessageBody(params.Get("Body"))
	if err != nil {
		log.Println("getting_measurements_from_message_failed: " + err.Error())
		return
	}

	err = s.sendMeasurementMessage(ctx, measurements, params.Get("From"))
	if err != nil {
		log.Println("failed to send measurement message: %w", err)
	}

	log.Println("sending message to queue succeeded")
}

func (s service) sendMeasurementMessage(ctx context.Context, measurements []models.Measurement, phoneNumber string) error {
	message := models.Message{
		Measurements: measurements,
		PhoneNumber:  phoneNumber,
		MessageType:  models.MessageTypeMeasurement,
	}

	return s.rabbitClient.Publish(ctx, "modules-messages", "", message)
}

func (s service) getMeasurementsFromMessageBody(body string) ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	for measurementType, regex := range parametersRegexsMatch {
		match := regex.FindAllStringSubmatch(body, -1)
		if len(match) == 0 || len(match[0]) < 2 {
			return nil, fmt.Errorf("no match in body for %s", measurementType)
		}

		parsedParameter, err := strconv.ParseFloat(match[0][1], 32)
		if err != nil {
			return nil, fmt.Errorf("invalid received parameter :%w", err)
		}

		measurements = append(measurements, models.Measurement{
			Name:  measurementType,
			Value: parsedParameter,
		})
	}

	return measurements, nil
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
