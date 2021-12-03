package middleware

import (
	"time"

	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AccessLogging(c *gin.Context) {
	start := time.Now()
	c.Next()
	log.Default().Info("access",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("latency", time.Now().Sub(start)),
	)
}
