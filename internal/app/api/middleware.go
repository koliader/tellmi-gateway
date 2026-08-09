package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

const traceIDHeader = "X-Trace-Id"

// requestLogger logs every request with its OTel trace context so Loki lines
// stay correlated with Tempo traces, and exposes the trace ID to the client.
func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ctx := c.Request.Context()
		sc := trace.SpanContextFromContext(ctx)
		if sc.IsValid() {
			c.Writer.Header().Set(traceIDHeader, sc.TraceID().String())
		}

		c.Next()

		log.Info().
			Ctx(ctx).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("duration", time.Since(start)).
			Msg("request")
	}
}
