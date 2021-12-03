package api

import (
	"net/http"
	"testing"

	"github.com/aurelius15/go-skeleton/internal/server"
	"github.com/steinfletcher/apitest"
)

func TestGetUser_Success(t *testing.T) {
	apitest.New().
		Handler(server.CreateServer("8888").Router).
		Get("/api/v1/users/1111").
		Expect(t).
		Body("Welcome!\n").
		Status(http.StatusOK).
		End()
}
