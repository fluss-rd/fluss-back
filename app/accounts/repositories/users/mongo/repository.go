package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/flussrd/fluss-back/app/accounts/models"
	repository "github.com/flussrd/fluss-back/app/accounts/repositories/users"
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

// GetUser returns a user from the database from a userID
func (repo MongoRepository) GetUser(ctx context.Context, userID string) (models.User, error) {
	usersCollection := repo.getUsersCollection()

	user := models.User{}

	err := usersCollection.FindOne(ctx, bson.M{
		"_id": userID,
	}).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.User{}, ErrNotFound
	}

	return models.User{}, nil
}

// SaveUser sabes a user on the database
func (repo MongoRepository) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	collection := repo.getUsersCollection()

	now := time.Now()

	user.CreationDate = now
	user.UpdateDate = now

	result, err := collection.InsertOne(ctx, user)
	if errors.As(err, &mongo.WriteException{}) {
		mongoErr, _ := err.(mongo.WriteException)
		switch mongoErr.WriteErrors[0].Code {
		case mongoDuplicateCode:
			return models.User{}, repository.ErrDuplicateFields
		}
	}

	fmt.Println(result.InsertedID)

	return user, nil
}

// AddRoleToUser adds a role to a user
func (repo MongoRepository) AddRoleToUser(ctx context.Context, userID string, role models.Role) error {
	usersCollection := repo.getUsersCollection()

	updateResult, err := usersCollection.UpdateOne(ctx, bson.M{
		"_id": userID,
	}, bson.M{
		"roleName": role,
	})

	if updateResult != nil && updateResult.MatchedCount == 0 {
		return ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user on the database
func (repo MongoRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (repo MongoRepository) getUsersCollection() *mongo.Collection {
	return repo.client.Database(databaseName).Collection("users")
}
