//go:build integration
// +build integration

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/aurelius15/go-skeleton/internal/entity"
	"github.com/aurelius15/go-skeleton/internal/helper"
	"github.com/aurelius15/go-skeleton/internal/server"
	"github.com/aurelius15/go-skeleton/internal/server/route"
	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/go-redis/redismock/v8"
	"github.com/steinfletcher/apitest"
)

func TestGetUser_Success(t *testing.T) {
	userID := helper.UUID()
	db, mock := redismock.NewClientMock()
	storage.SetInstance(db)

	testUser := entity.User{
		ID:        userID,
		FirstName: "test",
		LastName:  "test",
		Address:   nil,
	}

	expectedResponse, _ := json.Marshal(testUser)

	mock.ExpectGet(userID).SetVal(string(expectedResponse))

	apitest.New().
		Handler(server.NewServer()).
		Get(fmt.Sprintf("%s/%s", route.APIPrefix, userID)).
		Expect(t).
		Body(string(expectedResponse)).
		Status(http.StatusOK).
		End()
}
