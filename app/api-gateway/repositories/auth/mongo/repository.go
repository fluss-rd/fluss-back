package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/accounts/models"
	repository "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName = "flussDB"
)

type mongoRepository struct {
	client *mongo.Client
}

// New returns a new mongo repository
func New(client *mongo.Client) repository.Repository {
	return mongoRepository{
		client: client,
	}
}

// GetRole returns a new role from the role name
func (repo mongoRepository) GetRole(ctx context.Context, roleName string) (models.Role, error) {
	collection := repo.getRolesCollection()

	role := models.Role{}

	err := collection.FindOne(ctx, bson.M{
		"roleName": roleName,
	}).Decode(&role)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.Role{}, repository.ErrNotFound
	}

	return role, nil
}

func (repo mongoRepository) getRolesCollection() *mongo.Collection {
	return repo.client.Database(databaseName).Collection("roles")
}
