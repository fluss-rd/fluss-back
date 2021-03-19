package repository

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/accounts/models"
	repository "github.com/flussrd/fluss-back/accounts/repositories/roles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName = "flussDB"

	mongoDuplicateCode = 11000
)

var (
	ErrNotFound = errors.New("not found")
)

type mongoRepository struct {
	client mongo.Client
}

func New(ctx context.Context, client mongo.Client) mongoRepository {
	return mongoRepository{
		client: client,
	}
}

func (m mongoRepository) GetRole(ctx context.Context, roleName string) (models.Role, error) {
	rolesCollection := m.getRolesCollection()

	role := models.Role{}

	err := rolesCollection.FindOne(ctx, bson.M{
		"roleName": roleName,
	}).Decode(&role)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.Role{}, ErrNotFound
	}

	if err != nil {
		return models.Role{}, err
	}

	return models.Role{}, nil
}

func (m mongoRepository) CreateRole(ctx context.Context, role models.Role) error {
	collection := m.getRolesCollection()

	_, err := collection.InsertOne(ctx, role)
	if errors.As(err, &mongo.WriteException{}) {
		mongoErr, _ := err.(mongo.WriteException)
		switch mongoErr.WriteErrors[0].Code {
		case mongoDuplicateCode:
			return repository.ErrDuplicateFields
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) getRolesCollection() *mongo.Collection {
	return m.client.Database(databaseName).Collection("roles")
}
