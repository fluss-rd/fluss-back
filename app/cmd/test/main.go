package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	riverClient "github.com/flussrd/fluss-back/app/shared/grpc-clients/river-management"
)

func main() {
	os.Setenv("RIVER_MANAGEMENT_URL", "localhost:8080")
	fmt.Println(os.Getenv("RIVER_MANAGEMENT_URL"))
	riverserviceClient, err := riverClient.InitClient()
	if err != nil {
		log.Fatal("initializing river client failed", err.Error())
	}

	serviceClient := riverserviceClient.GetServiceClient()
	module, err := serviceClient.GetModuleByPhonenumber(context.Background(), &grpchandler.GetModuleRequest{PhoneNumber: "+18097538039"})
	if err != nil {
		log.Fatal("failed: ", err.Error())
	}

	fmt.Println(module)
}
