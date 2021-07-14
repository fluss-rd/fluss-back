package repository

import "context"

type Repository interface {
	GetData(ctx context.Context) error
}
