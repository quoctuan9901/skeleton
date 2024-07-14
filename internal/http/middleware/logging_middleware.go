package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggingMiddleware middleware ghi log cho request
func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		if statusCode >= 400 { // Chỉ ghi log Error level cho request lỗi
			logger.Error("Request error",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status_code", statusCode),
				zap.Duration("latency", latency),
			)
		} else {
			logger.Info("Request info",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status_code", statusCode),
				zap.Duration("latency", latency),
			)
		}
	}
}
