package users_client

import (
	"context"

	"github.com/koliader/tellmi-gateway/internal/domain/model"
	grpc_error "github.com/koliader/tellmi-gateway/internal/lib/error/grpc"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
	"github.com/koliader/tellmi-gateway/internal/pb"
	"google.golang.org/grpc/codes"
)

func (c *Client) Register(ctx context.Context, req model.RegisterReq) (*pb.AuthRes, *codes.Code, error) {
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := usersGrpcServiceClient.Register(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) Login(ctx context.Context, req model.LoginReq) (*pb.AuthRes, *codes.Code, error) {
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := usersGrpcServiceClient.Login(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) Refresh(ctx context.Context, req model.RefreshReq) (*pb.RefreshRes, *codes.Code, error) {
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.RefreshReq{
		RefreshToken: req.RefreshToken,
	}

	res, err := usersGrpcServiceClient.Refresh(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) GetUserById(ctx context.Context, req *model.IDReq) (*pb.UserRes, *codes.Code, error) {
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.IdReq{
		Id: req.ID,
	}

	res, err := usersGrpcServiceClient.GetUserById(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) ListUsers(ctx context.Context, headers model.AuthHeaders) (*pb.ListUserRes, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.Empty{}

	res, err := usersGrpcServiceClient.ListUsers(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) UpdateUser(ctx context.Context, req model.UpdateUserReq, headers model.AuthHeaders) (*pb.UserRes, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectUsersService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.UpdateUserReq{
		Id:       req.ID,
		Username: req.Username,
	}

	res, err := usersGrpcServiceClient.UpdateUser(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}
