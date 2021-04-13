package main

import (
	"context"
	"log"
	"net/http"

	"github.com/flussrd/fluss-back/accounts/config"
	handler "github.com/flussrd/fluss-back/accounts/handlers/http"
	rolesRepo "github.com/flussrd/fluss-back/accounts/repositories/roles/mongo"
	usersRepo "github.com/flussrd/fluss-back/accounts/repositories/users/mongo"
	"github.com/flussrd/fluss-back/accounts/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config, err := config.GetConfig("")
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

	rolesRepo := rolesRepo.New(client)
	usersRepo := usersRepo.New(client)

	service := service.NewService(usersRepo, rolesRepo)

	handler := handler.NewHTTPHandler(service)

	router := mux.NewRouter()

	router.Handle("/role", handler.HandleCreateRole(ctx)).Methods(http.MethodPost)

	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("err starting server: " + err.Error())
	}
}
