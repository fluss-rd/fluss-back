package client

import (
	"fmt"
	"os"

	"github.com/flussrd/fluss-back/app/river-management/handlers/grpc/grpchandler"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
}

// TODO: be able to pass options
func InitClient() (*Client, error) {
	fmt.Println(os.Getenv("RIVER_MANAGEMENT_URL"))
	conn, err := grpc.Dial(os.Getenv("RIVER_MANAGEMENT_URL"), grpc.WithInsecure()) // TODO: find out service discovery tools
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c Client) GetServiceClient() grpchandler.ServiceClient {
	return grpchandler.NewServiceClient(c.conn)
}
