package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
	"github.com/koliader/tellmi-sdk/model"
)

func (s *Server) createPost(ctx *gin.Context) {
	var req model.CreatePostReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*string)
	headers := model.AuthHeaders{
		Token: *authPayload,
	}

	res, code, err := s.postsClient.CreatePost(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) listPosts(ctx *gin.Context) {
	var req model.PaginationReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.postsClient.ListPosts(c, req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.Posts)
}

func (s *Server) getPostById(ctx *gin.Context) {
	var req model.IDReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.postsClient.GetPostByID(c, &req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) editPost(ctx *gin.Context) {
	var req model.EditPostReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*string)
	headers := model.AuthHeaders{
		Token: *authPayload,
	}

	res, code, err := s.postsClient.EditPost(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) deletePost(ctx *gin.Context) {
	var req model.IDReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*string)
	headers := model.AuthHeaders{
		Token: *authPayload,
	}

	res, code, err := s.postsClient.DeletePost(c, &req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
