package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	httpHandler "github.com/flussrd/fluss-back/app/reporting/handlers/http"
	influxRepository "github.com/flussrd/fluss-back/app/reporting/repositories/reports/influx"
	"github.com/flussrd/fluss-back/app/reporting/service"
	"github.com/gorilla/mux"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/subosito/gotenv"
)

var (
	defaultPort = "5000"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	fmt.Println(time.Now())
	ctx := context.Background()

	router := mux.NewRouter()

	port := os.Getenv("REST_PORT")
	if port == "" {
		port = defaultPort
	}

	influxClient := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))

	go func() { checkInfluxDB(ctx, influxClient) }()

	repo := influxRepository.New(influxClient)
	service := service.New(repo)
	handler := httpHandler.New(service, router)

	handler.HandleRoutes(ctx)

	log.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("could not start listening", err.Error())
	}
}

func checkInfluxDB(ctx context.Context, client influxdb2.Client) {
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

	if err != nil || !ok {
		log.Fatal("influx db not ready")
	}
}
