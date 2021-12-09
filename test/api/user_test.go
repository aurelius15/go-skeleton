//go:build integration
// +build integration

package api

import (
	"github.com/aurelius15/go-skeleton/internal/server"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func TestGetUser_Success(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	apitest.New().
		Handler(server.NewServer()).
		Get("/api/v1/users/1111").
		Expect(t).
		Body("Welcome!\n").
		Status(http.StatusOK).
		End()
}
