package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	s := NewServer()

	assert.NotNil(t, s)
}
