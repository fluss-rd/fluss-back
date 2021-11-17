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
	gorillaHandlers "github.com/gorilla/handlers"
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

	config, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatal("failed to load config: " + err.Error())
	}

	ctx := context.Background()

	fmt.Println(config.DatabaseConfig.Connection)

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

	router.Handle("/roles", handler.HandleCreateRole(ctx)).Methods(http.MethodPost)
	router.Handle("/roles", handler.HandleGetRoles(ctx)).Methods(http.MethodGet)

	router.Handle("/users", handler.HandleCreateUser(ctx)).Methods(http.MethodPost)
	router.Handle("/users", handler.HandleGetUsers(ctx)).Methods(http.MethodGet)
	router.Handle("/users/{id}", handler.HandleGetUser(ctx)).Methods(http.MethodGet)
	router.Handle("/users/{id}", handler.HandleUpdateUser(ctx)).Methods(http.MethodPatch)

	router.Handle("/login", handler.HandleLogin(ctx)).Methods(http.MethodPost)
	fmt.Println("Listening...")

	loggedRouter := gorillaHandlers.LoggingHandler(os.Stdout, router)

	err = http.ListenAndServe(":"+config.Port, loggedRouter)
	if err != nil {
		log.Fatal("err starting server: " + err.Error())
	}
}
