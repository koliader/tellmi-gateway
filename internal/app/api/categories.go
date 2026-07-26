package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koliader/tellmi-gateway/internal/domain/model"
	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
)

func (s *Server) createCategory(ctx *gin.Context) {
	var req model.CreateCategoryReq
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

	res, code, err := s.postsClient.CreateCategory(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) listCategories(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.postsClient.ListCategories(c)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.Categories)
}

func (s *Server) editCategory(ctx *gin.Context) {
	var req model.EditCategoryReq
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

	res, code, err := s.postsClient.EditCategory(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
