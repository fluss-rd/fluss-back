package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/river-management/models"
	repository "github.com/flussrd/fluss-back/app/river-management/repositories/modules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName = "flussDB"
)

type mongoRepository struct {
	client *mongo.Client
}

// New returns a new mongoRepository instance
func New(client *mongo.Client) repository.Repository {
	return mongoRepository{client: client}
}

func (r mongoRepository) GetModule(ctx context.Context, moduleID string) (models.Module, error) {
	modulesCollection := r.getModulesCollection()

	module := models.Module{}

	err := modulesCollection.FindOne(ctx, bson.M{
		"_id": moduleID,
	}).Decode(&module)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.Module{}, repository.ErrNotFound
	}

	if err != nil {
		return models.Module{}, err
	}

	return models.Module{}, nil
}

func (r mongoRepository) GetAllModulesWithOutPagination(ctx context.Context) ([]models.Module, error) {
	return r.getModulesWithoutPagination(ctx, bson.M{})
}

func (r mongoRepository) GetAllModules(ctx context.Context) ([]models.Module, string, error) {
	return nil, "", nil
}

func (r mongoRepository) GetModulesByRiverWithoutPagination(ctx context.Context, riverID string) ([]models.Module, error) {
	return r.getModulesWithoutPagination(ctx, bson.M{"riverID": riverID})
}

func (r mongoRepository) getModulesWithoutPagination(ctx context.Context, filter bson.M) ([]models.Module, error) {
	modulesCollection := r.getModulesCollection()

	modules := []models.Module{}

	result, err := modulesCollection.Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository.ErrNotFound
	}

	err = result.All(ctx, &modules)
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (r mongoRepository) GetModulesByRiver(ctx context.Context) ([]models.Module, string, error) {
	return nil, "", nil
}

func (r mongoRepository) SaveModule(ctx context.Context, module models.Module) error {
	return nil
}

func (r mongoRepository) getModulesCollection() *mongo.Collection {
	return r.client.Database(databaseName).Collection("modules")
}