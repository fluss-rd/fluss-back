package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flussrd/fluss-back/app/accounts/config"
	repository "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth/mongo"
	"github.com/flussrd/fluss-back/app/api-gateway/router"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var endpoints = []router.Endpoint{
	{
		Path:       "/account/login",
		RemotePath: "/login",
		RemotHost:  "http://accounts:5000",
		Method:     http.MethodPost,
		Authorized: false,
	},
	{
		Path:       "/account/roles",
		RemotePath: "/role",
		RemotHost:  "http://accounts:5000",
		Method:     http.MethodPost,
		Authorized: true,
	},
	{
		Path:       "/account/roles",
		RemotePath: "/role",
		RemotHost:  "http://accounts:5000",
		Method:     http.MethodGet,
		Authorized: true,
	},
	{
		Path:       "/account/users",
		RemotePath: "/role",
		RemotHost:  "http://accounts:5000",
		Method:     http.MethodPost,
		Authorized: true,
	},
}

// func init() {
// 	_ = gotenv.Load()
// }

func main() {
	config, err := config.GetConfig(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Fatal("could not load config: " + err.Error())
	}

	ctx := context.Background()

	handler := mux.NewRouter()

	client, err := getMongoClient(ctx, config.DatabaseConfig.Connection)
	if err != nil {
		log.Fatal("err creating mongo client: " + err.Error())
	}

	go func() {
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(fmt.Errorf("pinging database failed: %w", err))
		}
	}()

	repo := repository.New(client)

	_, err = router.NewRouter(ctx, endpoints, repo, handler)
	if err != nil {
		log.Fatal("could not create router: " + err.Error())
	}

	handler.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	loggedRouter := gorillaHandlers.LoggingHandler(os.Stdout, handler)

	fmt.Println("Listening...")
	err = http.ListenAndServe(":5000", loggedRouter)
	if err != nil {
		log.Fatal("listeting_starting_failed: " + err.Error())
	}
}

func getMongoClient(ctx context.Context, connectionURL string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func handleIndex(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Hello world!"))
}
