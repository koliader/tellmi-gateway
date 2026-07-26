package posts_client

import (
	"context"

	"github.com/koliader/tellmi-gateway/internal/domain/model"
	grpc_error "github.com/koliader/tellmi-gateway/internal/lib/error/grpc"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
	"github.com/koliader/tellmi-gateway/internal/pb"
	"google.golang.org/grpc/codes"
)

func (c *Client) CreatePost(ctx context.Context, req model.CreatePostReq, headers model.AuthHeaders) (*pb.Post, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.CreatePostReq{
		Title:      req.Title,
		Description: req.Description,
		CategoryId: req.CategoryID,
	}

	res, err := postsGrpcServiceClient.CreatePost(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) ListPosts(ctx context.Context) (*pb.ListPostsRes, *codes.Code, error) {
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.Empty{}

	res, err := postsGrpcServiceClient.ListPosts(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) GetPostByID(ctx context.Context, req *model.IDReq) (*pb.PostRow, *codes.Code, error) {
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.GetByIDReq{
		Id: req.ID,
	}

	res, err := postsGrpcServiceClient.GetPostByID(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) EditPost(ctx context.Context, req model.EditPostReq, headers model.AuthHeaders) (*pb.Post, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.EditPostReq{
		Id:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		CategoryId:  req.CategoryID,
	}

	res, err := postsGrpcServiceClient.EditPost(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) DeletePost(ctx context.Context, req *model.IDReq, headers model.AuthHeaders) (*pb.Success, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.GetByIDReq{
		Id: req.ID,
	}

	res, err := postsGrpcServiceClient.DeletePost(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

// Categories

func (c *Client) CreateCategory(ctx context.Context, req model.CreateCategoryReq, headers model.AuthHeaders) (*pb.Category, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.CreateCategoryReq{
		Name: req.Name,
	}

	res, err := postsGrpcServiceClient.CreateCategory(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) ListCategories(ctx context.Context) (*pb.ListCategoriesRes, *codes.Code, error) {
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.Empty{}

	res, err := postsGrpcServiceClient.ListCategories(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) EditCategory(ctx context.Context, req model.EditCategoryReq, headers model.AuthHeaders) (*pb.Success, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.EditCategoryReq{
		Id:   req.ID,
		Name: req.Name,
	}

	res, err := postsGrpcServiceClient.EditCategory(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

// Comments

func (c *Client) CreateComment(ctx context.Context, req model.CreateCommentReq, headers model.AuthHeaders) (*pb.Comment, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.CreateCommentReq{
		Comment: req.Comment,
		PostId:  req.PostID,
	}

	res, err := postsGrpcServiceClient.CreateComment(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) ListCommentsByPost(ctx context.Context, req *model.IDReq) (*pb.ListCommentsRes, *codes.Code, error) {
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.GetByIDReq{
		Id: req.ID,
	}

	res, err := postsGrpcServiceClient.ListCommentsByPost(ctx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) EditComment(ctx context.Context, req model.EditCommentReq, headers model.AuthHeaders) (*pb.Comment, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.EditCommentReq{
		Id:      req.ID,
		Comment: req.Comment,
	}

	res, err := postsGrpcServiceClient.EditComment(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}

func (c *Client) DeleteComment(ctx context.Context, req *model.IDReq, headers model.AuthHeaders) (*pb.Success, *codes.Code, error) {
	authCtx := middleware.CreateAuthMetadata(ctx, headers.Token)
	if err := c.ConnectPostsService(&ctx); err != nil {
		return nil, nil, err
	}

	arg := pb.GetByIDReq{
		Id: req.ID,
	}

	res, err := postsGrpcServiceClient.DeleteComment(authCtx, &arg)
	if err != nil {
		return nil, grpc_error.GetErrorCode(err), grpc_error.ErrorResponse(err)
	}

	return res, nil, nil
}
