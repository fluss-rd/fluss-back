package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	riverClient "github.com/flussrd/fluss-back/app/shared/grpc-clients/river-management"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
	httpHandler "github.com/flussrd/fluss-back/app/telemetry/handlers/http"
	rabbitHandler "github.com/flussrd/fluss-back/app/telemetry/handlers/rabbitmq"
	repository "github.com/flussrd/fluss-back/app/telemetry/repositories/measurements/influx"
	"github.com/flussrd/fluss-back/app/telemetry/service"
	"github.com/gorilla/mux"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	restPort = os.Getenv("REST_PORT")
)

func init() {
	if restPort == "" {
		restPort = "5000"
	}
}

func main() {
	ctx := context.Background()
	influxDbURL := os.Getenv("INFLUXDB_URL")
	influxDbToken := os.Getenv("INFLUXDB_TOKEN")

	client := influxdb2.NewClient(influxDbURL, influxDbToken)

	defer client.Close()

	var err error
	var ok bool

	for i := 0; i < 10; i++ {
		ok, err = client.Ready(ctx)
		if ok {
			break
		}

		if err != nil || !ok {
			time.Sleep(time.Second * 2)
		}
	}

	if err != nil {
		log.Fatal("influxdb not ready: ", err.Error())
	}

	repo := repository.New(client)

	// var grpcConn *grpc.ClientConn
	// conn, err :=

	riverserviceClient, err := riverClient.InitClient()
	if err != nil {
		log.Fatal("initializing river client failed", err.Error())
	}

	var rabbitClient rabbit.RabbitClient
	for i := 0; i < 10; i++ {
		rabbitClient, err = rabbit.InitRabbitClient(os.Getenv("RABBITMQ_URL"))

		if err == nil {
			break
		}

		time.Sleep(time.Second * 2)
	}

	if err != nil {
		log.Fatal("initializing rabbit client failed", err.Error())
	}

	service := service.New(riverserviceClient, repo, rabbitClient)

	handlers := rabbitHandler.NewRabbitHandler(rabbitClient, service)

	go func() {
		router := mux.NewRouter()

		httpHandler := httpHandler.NewHTTPHandler(service, router)

		httpHandler.Init(ctx)

		log.Println("listening for messages at rest port...")

		err = http.ListenAndServe(":"+restPort, router)
		if err != nil {
			log.Fatal("could not start to listen at rest port: " + err.Error())
		}
	}()

	go func() {
		err = handlers.HandleMessages(ctx)
		if err != nil {
			log.Fatal("could not start to listen to messages", err.Error())
		}
	}()

	forever := make(chan bool)
	<-forever
}
