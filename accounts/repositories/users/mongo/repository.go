package repository

import (
	"context"

	"github.com/flussrd/fluss-back/accounts/models"
	"go.mongodb.org/mongo-driver/mongo"
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

// GetUser returns a user from the database from a userID
func (repo MongoRepository) GetUser(ctx context.Context, userID string) (models.User, error) {
	return models.User{}, nil
}

// SaveUser sabes a user on the database
func (repo MongoRepository) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

// AddRoleToUser adds a role to a user
func (repo MongoRepository) AddRoleToUser(ctx context.Context, userID string, role models.Role) error {
	return nil
}

// UpdateUser updates a user on the database
func (repo MongoRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}
