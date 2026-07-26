package posts_client

import (
	"context"
	"fmt"

	"github.com/koliader/tellmi-gateway/internal/config"
	"github.com/koliader/tellmi-gateway/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var postsGrpcServiceClient pb.PostsClient

type Client struct {
	config config.Config
}

func NewClient(config config.Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) ConnectPostsService(ctx *context.Context) error {
	conn, err := grpc.DialContext(*ctx, c.config.PostsServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		postsGrpcServiceClient = nil
		return fmt.Errorf("connection to posts gRPC service failed: %v", err)
	}
	if postsGrpcServiceClient != nil {
		conn.Close()
		return nil
	}
	postsGrpcServiceClient = pb.NewPostsClient(conn)
	return nil
}
