package users_client

import (
	"context"
	"fmt"

	"github.com/koliader/tellmi-gateway/internal/config"
	"github.com/koliader/tellmi-gateway/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var usersGrpcServiceClient pb.UsersClient

type Client struct {
	config config.Config
}

func NewClient(config config.Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) ConnectUsersService(ctx *context.Context) error {
	conn, err := grpc.DialContext(*ctx, c.config.UsersServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		usersGrpcServiceClient = nil
		return fmt.Errorf("connection to users gRPC service failed: %v", err)
	}
	if usersGrpcServiceClient != nil {
		conn.Close()
		return nil
	}
	usersGrpcServiceClient = pb.NewUsersClient(conn)
	return nil
}
