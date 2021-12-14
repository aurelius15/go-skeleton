package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCorrelation_Set_Header(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Correlation(c)
	assert.NotEmpty(t, c.GetHeader(CorrelationHeader))
}

func TestCorrelation_Do_Not_Overwrite(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)
	request.Header.Set(CorrelationHeader, "test")

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Correlation(c)
	assert.Equal(t, c.GetHeader(CorrelationHeader), "test")
}
