package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flussrd/fluss-back/app/accounts/config"
	handler "github.com/flussrd/fluss-back/app/accounts/handlers/http"
	rolesRepo "github.com/flussrd/fluss-back/app/accounts/repositories/roles/mongo"
	usersRepo "github.com/flussrd/fluss-back/app/accounts/repositories/users/mongo"
	"github.com/flussrd/fluss-back/app/accounts/service"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	_ = gotenv.Load()
}

func main() {

	config, err := config.GetConfig(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Fatal("failed to load config: " + err.Error())
	}

	ctx := context.Background()

	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseConfig.Connection))
	if err != nil {
		log.Fatal("err creating mongo client: " + err.Error())
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("err connecting to database: " + err.Error())
	}

	go func() {
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(fmt.Errorf("pinging database failed: %w", err))
		}
	}()

	rolesRepo := rolesRepo.New(client)
	usersRepo := usersRepo.New(client)

	service := service.NewService(usersRepo, rolesRepo)

	handler := handler.NewHTTPHandler(service)

	router := mux.NewRouter()

	router.Handle("/role", handler.HandleCreateRole(ctx)).Methods(http.MethodPost)
	router.Handle("/role", handler.HandleGetRoles(ctx)).Methods(http.MethodGet)

	router.Handle("/user", handler.HandleCreateUser(ctx)).Methods(http.MethodPost)
	fmt.Println("Listening...")

	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("err starting server: " + err.Error())
	}
}
