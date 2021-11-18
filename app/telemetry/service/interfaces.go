package service

import (
	"context"

	"github.com/flussrd/fluss-back/app/telemetry/models"
)

type Service interface {
	SaveMeasurement(ctx context.Context, message models.Message) error
	HandleHTTPMessage(ctx context.Context, source string, body string)
}
