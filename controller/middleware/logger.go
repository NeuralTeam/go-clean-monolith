package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"go-clean-monolith/pkg/logger"
	"net/http"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------

// LoggerMiddleware description
type LoggerMiddleware struct {
	logger logger.Logger
}

// NewLoggerMiddleware description
func NewLoggerMiddleware(logger logger.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		logger: logger,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Setup description
func (m LoggerMiddleware) Setup() gin.HandlerFunc {
	type loggerFormatter struct {
		Request    *http.Request
		TimeStamp  time.Time
		StatusCode int
		Method     string
		Path       string
		Latency    time.Duration
		ClientIP   string
		BodySize   int
		UserID     int64
	}

	return func(ctx *gin.Context) {
		// Before request
		requestId := uuid.NewString()

		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Set("request_id", requestId)

		// Process request
		ctx.Next()

		// After request
		params := loggerFormatter{
			Request: ctx.Request,
		}

		// Stop timer
		params.TimeStamp = time.Now()
		params.Latency = params.TimeStamp.Sub(start)

		params.StatusCode = ctx.Writer.Status()
		params.Method = ctx.Request.Method
		params.ClientIP = ctx.ClientIP()
		params.BodySize = ctx.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		params.Path = path

		if userId, ok := ctx.Get("user_id"); ok {
			params.UserID = userId.(int64)
		} else {
			params.UserID = -1
		}

		if params.StatusCode == 500 {
			st, _ := ctx.Get("stack_trace")
			b, _ := json.Marshal(st.(Stacks))
			m.logger.Trace().
				Str("request_id", requestId).
				Int("status", params.StatusCode).
				Str("method", params.Method).
				Str("path", params.Path).
				Dur("latency", params.Latency).
				Int("size", params.BodySize).
				Str("addr", params.ClientIP).
				Int64("user_id", params.UserID).
				RawJSON("stack", b).
				Send()
		} else {
			m.logger.Info().
				Str("request_id", requestId).
				Int("status", params.StatusCode).
				Str("method", params.Method).
				Str("path", params.Path).
				Dur("latency", params.Latency).
				Int("size", params.BodySize).
				Str("addr", params.ClientIP).
				Int64("user_id", params.UserID).
				Send()
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
