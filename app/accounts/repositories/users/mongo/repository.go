package repository

import (
	"context"
	"errors"
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
	return repo.getUser(ctx, bson.M{"_id": userID})
}

// GetUserByEmail returns a use from the database with the specified email
func (repo MongoRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return repo.getUser(ctx, bson.M{"email": email})
}

func (repo MongoRepository) getUser(ctx context.Context, filter bson.M) (models.User, error) {
	usersCollection := repo.getUsersCollection()

	user := models.User{}

	err := usersCollection.FindOne(ctx, filter).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.User{}, repository.ErrNotFound
	}

	return user, nil
}

// SaveUser sabes a user on the database
func (repo MongoRepository) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	collection := repo.getUsersCollection()

	now := time.Now()

	user.CreationDate = now
	user.UpdateDate = now

	_, err := collection.InsertOne(ctx, user)
	if errors.As(err, &mongo.WriteException{}) {
		mongoErr, _ := err.(mongo.WriteException)
		switch mongoErr.WriteErrors[0].Code {
		case mongoDuplicateCode:
			return models.User{}, repository.ErrDuplicateFields{}
		}
	}

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
		return repository.ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user on the database
func (repo MongoRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	updateFields := bson.M{}

	// TODO: research on standards on how to do this.
	// My opinion is that the database layer should not be aware of which fields should be changed,
	// as this might be a business requirement
	if user.UserID == "" {
		return models.User{}, repository.ErrMissingUserID
	}

	if user.Email != "" {
		updateFields["email"] = user.Email
	}

	if user.Password != "" {
		updateFields["password"] = user.Password
	}

	if len(updateFields) == 0 {
		return models.User{}, repository.ErrNothingToUpdate
	}

	usersCollection := repo.getUsersCollection()

	_, err := usersCollection.UpdateOne(ctx, bson.M{"_id": user.UserID}, bson.M{
		"$set": updateFields,
	})

	// TODO: handle when id does not exits / no matches
	if err != nil {
		// TODO: wrap this
		return models.User{}, err
	}

	// TODO: see how to retrieve the changed params
	return models.User{}, nil
}

func (repo MongoRepository) getUsersCollection() *mongo.Collection {
	return repo.client.Database(databaseName).Collection("users")
}
