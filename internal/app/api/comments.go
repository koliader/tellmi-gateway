package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koliader/tellmi-sdk/model"
	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
)

func (s *Server) createComment(ctx *gin.Context) {
	var req model.CreateCommentReq
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

	res, code, err := s.postsClient.CreateComment(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) listCommentsByPost(ctx *gin.Context) {
	var req model.IDReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.postsClient.ListCommentsByPost(c, &req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.Comments)
}

func (s *Server) editComment(ctx *gin.Context) {
	var req model.EditCommentReq
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

	res, code, err := s.postsClient.EditComment(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) deleteComment(ctx *gin.Context) {
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

	res, code, err := s.postsClient.DeleteComment(c, &req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
