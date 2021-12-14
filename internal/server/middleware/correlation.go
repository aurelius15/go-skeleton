package middleware

import (
	"github.com/aurelius15/go-skeleton/internal/helper"
	"github.com/gin-gonic/gin"
)

const CorrelationHeader = "X-Correlation-ID"

func Correlation(c *gin.Context) {
	if c.GetHeader(CorrelationHeader) == "" {
		c.Request.Header.Set(CorrelationHeader, helper.UUID())
	}
}
