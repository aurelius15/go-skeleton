package server

import (
	"github.com/aurelius15/go-skeleton/internal/server/route"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	route.Routes = []route.Route{
		route.Route{
			Method: "GET",
			Path:   "my-prefix",
			Handle: func(context *gin.Context) {},
		},
	}

	server := CreateServer("9000")

	assert.Equal(t, server.port, ":9000")
	assert.Len(t, server.Router.Routes(), 1)
}
