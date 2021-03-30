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
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// MongoRepository repository entitie to be used with mongoDB
type MongoRepository struct {
	client *mongo.Client
}

// New returns a new mongo repository
func New(client *mongo.Client) MongoRepository {
	return MongoRepository{
		client: client,
	}
}

// GetRole returns a role from the database from the roleName
func (m MongoRepository) GetRole(ctx context.Context, roleName string) (models.Role, error) {
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

// CreateRole stores a new role in the database
func (m MongoRepository) CreateRole(ctx context.Context, role models.Role) error {
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

func (m MongoRepository) getRolesCollection() *mongo.Collection {
	return m.client.Database(databaseName).Collection("roles")
}
