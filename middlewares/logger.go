package middlewares

import (
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/plusone/utils/logger"
)

// LoggerMiddleware 是一个Gin中间件，用于结构化日志记录和请求追踪
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 为每个请求生成一个唯一的 Trace ID
		traceID := uuid.NewString()

		// 创建一个带有 Trace ID 的 logger
		slogWithTrace := slog.Default().With("trace_id", traceID)

		// 将 Trace ID 和 logger 存入 context
		ctx := logger.WithTraceID(c.Request.Context(), traceID)
		ctx = logger.WithLogger(ctx, slogWithTrace)

		// 更新 gin.Context 中的 Request
		c.Request = c.Request.WithContext(ctx)

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)

		// 记录请求日志
		slogWithTrace.Info("request handled",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency.String(),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
