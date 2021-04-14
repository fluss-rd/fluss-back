package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/flussrd/fluss-back/app/accounts/models"
	repository "github.com/flussrd/fluss-back/app/accounts/repositories/roles"
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
type mongoRepository struct {
	client *mongo.Client
}

// New returns a new mongo repository
func New(client *mongo.Client) repository.Repository {
	return mongoRepository{
		client: client,
	}
}

// GetRole returns a role from the database from the roleName
func (m mongoRepository) GetRole(ctx context.Context, roleName string) (models.Role, error) {
	rolesCollection := m.getRolesCollection()

	role := models.Role{}

	err := rolesCollection.FindOne(ctx, bson.M{
		"roleName": roleName,
	}).Decode(&role)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.Role{}, repository.ErrNotFound
	}

	if err != nil {
		return models.Role{}, err
	}

	return models.Role{}, nil
}

// CreateRole stores a new role in the database
func (m mongoRepository) CreateRole(ctx context.Context, role models.Role) error {
	collection := m.getRolesCollection()

	result, err := collection.InsertOne(ctx, role)
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

	fmt.Println("inserted id:")
	fmt.Println(result.InsertedID)

	return nil
}

func (m mongoRepository) GetRoles(ctx context.Context) ([]models.Role, error) {
	collection := m.getRolesCollection()

	results, err := collection.Find(ctx, bson.D{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNotFound
	}

	// TODO: wrap this error
	if err != nil {
		return nil, err
	}

	roles := []models.Role{}

	err = results.All(ctx, &roles)
	if err != nil { //TODO: wrap
		return nil, err
	}

	return roles, nil
}

func (m mongoRepository) GetUserRole(ctx context.Context, userID string) ([]models.Role, error) {
	return nil, nil
}

func (m mongoRepository) getRolesCollection() *mongo.Collection {
	return m.client.Database(databaseName).Collection("roles")
}
