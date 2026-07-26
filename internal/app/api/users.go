package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koliader/tellmi-gateway/internal/domain/model"
	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
)

func (s *Server) register(ctx *gin.Context) {
	var req model.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.usersClient.Register(c, req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) login(ctx *gin.Context) {
	var req model.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.usersClient.Login(c, req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) getUserById(ctx *gin.Context) {
	var req model.IDReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api_error.ErrorInvalidArguments(err))
		return
	}

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, code, err := s.usersClient.GetUserById(c, &req)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.User)
}

func (s *Server) listUsers(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*string)
	headers := model.AuthHeaders{
		Token: *authPayload,
	}

	res, code, err := s.usersClient.ListUsers(c, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.Users)
}

func (s *Server) updateUser(ctx *gin.Context) {
	var req model.UpdateUserReq
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

	res, code, err := s.usersClient.UpdateUser(c, req, headers)
	if err != nil {
		ctx.JSON(api_error.ErrorCode(code), api_error.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res.User)
}
