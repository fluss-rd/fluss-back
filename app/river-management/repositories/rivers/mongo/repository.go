package repository

import (
	"context"
	"errors"
	"time"

	"github.com/flussrd/fluss-back/app/river-management/models"
	repository "github.com/flussrd/fluss-back/app/river-management/repositories/rivers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName = "fluss-db"

	mongoDuplicateCode = 11000
)

type mongoRepository struct {
	client *mongo.Client
}

// New returns a new mongo repository instance
func New(client *mongo.Client) repository.Repository {
	return mongoRepository{
		client: client,
	}
}

// SaveRiver saves a river on the database
func (r mongoRepository) SaveRiver(ctx context.Context, river models.River) (models.River, error) {
	riversCollection := r.getRiversCollection()

	river.CreationDate = time.Now()
	river.UpdateDate = time.Now()

	_, err := riversCollection.InsertOne(ctx, river)
	if errors.As(err, &mongo.WriteException{}) {
		mongoErr, _ := err.(mongo.WriteException)
		switch mongoErr.WriteErrors[0].Code {
		case mongoDuplicateCode:
			return models.River{}, repository.ErrDuplicateFields
		}
	}

	if err != nil {
		return models.River{}, err
	}

	return river, nil
}

// GetRiver finds a river on the database based in the id
func (r mongoRepository) GetRiver(ctx context.Context, riverID string) (models.River, error) {
	riversCollection := r.getRiversCollection()

	river := models.River{}

	err := riversCollection.FindOne(ctx, bson.M{
		"_id": riverID,
	}).Decode(&river)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.River{}, repository.ErrNotFound
	}

	return river, nil
}

// GetAllRiversNotPaginated temporal method for retrieving all the rivers from the database
func (r mongoRepository) GetAllRiversNotPaginated(ctx context.Context) ([]models.River, error) {
	riversCollection := r.getRiversCollection()

	results, err := riversCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	rivers := []models.River{}

	err = results.All(ctx, &rivers)
	if err != nil {
		return nil, err
	}

	return rivers, nil
}

// GetAllRivers returns all rivers from the database, paginated
func (r mongoRepository) GetAllRivers(ctx context.Context) ([]models.River, string, error) {
	return nil, "", nil
}

func (r mongoRepository) getRiversCollection() *mongo.Collection {
	return r.client.Database(databaseName).Collection("rivers")
}
