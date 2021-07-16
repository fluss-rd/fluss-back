package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/flussrd/fluss-back/app/accounts/config"
	grpcHandlers "github.com/flussrd/fluss-back/app/river-management/handlers/grpc"
	grpcCode "github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	handlers "github.com/flussrd/fluss-back/app/river-management/handlers/http"
	modulesRepository "github.com/flussrd/fluss-back/app/river-management/repositories/modules/mongo"
	riversRepository "github.com/flussrd/fluss-back/app/river-management/repositories/rivers/mongo"
	"github.com/flussrd/fluss-back/app/river-management/service"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
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

	client, err := getMongoClient(ctx, config.DatabaseConfig.Connection)
	if err != nil {
		log.Fatal("failed to get client: " + err.Error())
	}

	go func() {
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal("pinging database failed: " + err.Error())
		}
	}()

	riversRepo := riversRepository.New(client)
	modulesRepo := modulesRepository.New(client)

	service := service.New(riversRepo, modulesRepo)

	handler := handlers.NewHTTPHandler(service)

	router := mux.NewRouter()

	router.Handle("/rivers", handler.HandleGetRivers(ctx)).Methods(http.MethodGet)
	router.Handle("/rivers/{id}", handler.HandleGetRiver(ctx)).Methods(http.MethodGet)
	router.Handle("/rivers", handler.HandleCreateRiver(ctx)).Methods(http.MethodPost)

	router.Handle("/modules", handler.HandleGetModules(ctx)).Methods(http.MethodGet)
	router.Handle("/modules/{id}", handler.HandleGetModule(ctx)).Methods(http.MethodGet)
	router.Handle("/modules/{id}", handler.HandleUpdateModule(ctx)).Methods(http.MethodPatch)
	router.Handle("/modules", handler.HandleCreateModule(ctx)).Methods(http.MethodPost)

	fmt.Println("Listening on port " + config.Port)

	grpcHandler := grpcHandlers.NewHandler(service)

	grpcServer := grpc.NewServer()
	grpcCode.RegisterServiceServer(grpcServer, grpcHandler)

	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to start listener")
	}

	forever := make(chan bool)

	go func() {
		err = http.ListenAndServe(":"+config.Port, router)
		if err != nil {
			log.Fatal("failed to start listening http")
		}
	}()

	go func() {
		err = grpcServer.Serve(grpcListener)
		if err != nil {
			log.Fatal("failed to start listening grpc server")
		}
	}()

	fmt.Println("succesfully started server")

	<-forever
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
