package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
)

func (m *Middleware) AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, token, err := authHandleFunc(ctx, m.tokenMaker)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api_error.ErrorResponse(err))
			return
		}
		if payload.Role != "ADMIN" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				api_error.ErrorResponse(fmt.Errorf("no access")),
			)
			return
		}
		ctx.Set(AuthorizationPayloadKey, token)
		ctx.Next()
	}
}
