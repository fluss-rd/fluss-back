package client

import (
	grpchandler "github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	"google.golang.org/grpc"
)

// TODO: be able to pass options
func InitClient() (*grpchandler.ServiceClient, error) {
	conn, err := grpc.Dial("river-management:5000", grpc.WithInsecure()) // TODO: find out service discovery tools
	if err != nil {
		return nil, err
	}

	client := grpchandler.NewServiceClient(conn)

	return &client, nil
}
