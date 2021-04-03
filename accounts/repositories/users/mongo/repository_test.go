package repository

import (
	"context"
	"testing"

	"github.com/flussrd/fluss-back/accounts/models"
	"github.com/stretchr/testify/require"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSaveUser(t *testing.T) {
	c := require.New(t)

	mongoServer, err := strikememongo.Start("4.0.5")
	c.Nil(err)

	defer mongoServer.Stop()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoServer.URI()))
	c.Nil(err)

	ctx := context.Background()

	err = client.Connect(ctx)
	c.Nil(err)

	defer func() {
		err := client.Disconnect(ctx)
		c.Nil(err)
	}()

	repo := New(client)

	user := models.User{
		UserID:      "USR1234",
		PhoneNumber: "+18091111111",
	}

	insertedUser, err := repo.SaveUser(ctx, user)
	c.Nil(err)
	c.NotNil(insertedUser.CreationDate)
	c.NotNil(insertedUser.UpdateDate)
}
