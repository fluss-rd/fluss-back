package main

import (
	"context"
	"log"
	"os"
	"time"

	riverClient "github.com/flussrd/fluss-back/app/shared/grpc-clients/river-management"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
	rabbitHandler "github.com/flussrd/fluss-back/app/telemetry/handlers/rabbitmq"
	repository "github.com/flussrd/fluss-back/app/telemetry/repositories/measurements/influx"
	"github.com/flussrd/fluss-back/app/telemetry/service"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// func init() {
// 	gotenv.Load()
// }

func main() {
	ctx := context.Background()

	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))

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

	service := service.New(riverserviceClient, repo)

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

	handlers := rabbitHandler.NewRabbitHandler(rabbitClient, service)

	go func() {
		err = handlers.HandleMessages(ctx)
		if err != nil {
			log.Fatal("could not start to listen to messages", err.Error())
		}
	}()

	forever := make(chan bool)
	<-forever
}
