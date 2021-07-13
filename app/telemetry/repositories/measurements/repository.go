package repository

import (
	"context"

	"github.com/flussrd/fluss-back/app/telemetry/models"
)

type Repository interface {
	SaveMeasurement(ctx context.Context, measurement models.Message) error
}
