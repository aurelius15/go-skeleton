package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/aurelius15/go-skeleton/internal/helper"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestAccessLogging(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request

	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	log.SetDefault(observedLogger)

	AccessLogging(c)

	assert.Equal(t, observedLogs.Len(), 1)
	assert.Equal(t, observedLogs.All()[0].Message, "access")

	fields := observedLogs.All()[0].ContextMap()

	assert.Contains(t, fields, "method")
	assert.Contains(t, fields, "path")
	assert.Contains(t, fields, "status")
	assert.Contains(t, fields, "uuid")
	assert.Contains(t, fields, "latency")

	assert.Equal(t, fields["method"], "GET")
	assert.Equal(t, fields["path"], "/test")
	assert.Equal(t, fields["status"], int64(200))
	assert.Regexp(t, fields["uuid"], helper.UUIDMask())
}
