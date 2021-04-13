package repository

import (
	"context"
	"testing"

	"github.com/flussrd/fluss-back/app/accounts/models"
	"github.com/stretchr/testify/require"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateRole(t *testing.T) {
	c := require.New(t)

	mongoServer, err := strikememongo.Start("4.0.5")
	c.Nil(err)

	defer mongoServer.Stop()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoServer.URI()))
	c.Nil(err)

	ctx := context.Background()

	err = client.Connect(ctx)
	c.Nil(err)

	repo := mongoRepository{client: client}

	err = repo.CreateRole(context.Background(), models.Role{Name: "Admin"})
	c.Nil(err)

	defer func() {
		err := repo.client.Disconnect(ctx)
		c.Nil(err)
	}()
}
