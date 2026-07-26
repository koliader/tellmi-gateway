package api_error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ErrorInvalidArguments(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ErrorCode(code *codes.Code) int {
	httpCode := http.StatusInternalServerError
	if code != nil {
		switch *code {
		case codes.AlreadyExists:
			httpCode = http.StatusBadRequest
		case codes.NotFound:
			httpCode = http.StatusNotFound
		case codes.Unauthenticated:
			httpCode = http.StatusUnauthorized
		case codes.PermissionDenied:
			httpCode = http.StatusUnauthorized
		case codes.InvalidArgument:
			httpCode = http.StatusBadRequest
		case codes.Internal:
			httpCode = http.StatusInternalServerError
		}
	}
	return httpCode
}
