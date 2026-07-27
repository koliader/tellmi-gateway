package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	api_error "github.com/koliader/tellmi-gateway/internal/lib/error/api"
	"github.com/koliader/tellmi-sdk/token"
)

const (
	authHeaderKey           = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func authHandleFunc(ctx *gin.Context, tokenMaker token.Maker) (*token.Payload, *string, error) {
	authHeader := ctx.GetHeader(authHeaderKey)
	if len(authHeader) == 0 {
		err := errors.New("authorization header not specified")
		return nil, nil, err
	}
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return nil, nil, err
	}
	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %v", authorizationType)
		return nil, nil, err
	}
	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, nil, err
	}
	return payload, &accessToken, nil
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, token, err := authHandleFunc(ctx, m.tokenMaker)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api_error.ErrorResponse(err))
			return
		}
		ctx.Set(AuthorizationPayloadKey, token)
		ctx.Next()
	}
}
