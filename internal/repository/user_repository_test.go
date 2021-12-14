package repository

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aurelius15/go-skeleton/internal/entity"
	"github.com/aurelius15/go-skeleton/internal/helper"
	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {
	userRepositoryInstance = nil
}

// Get a user that exists
func (s *TestSuite) TestUserRepository_GetUserByID_Success() {
	userID := helper.UUID()
	db, mock := redismock.NewClientMock()
	storage.SetInstance(db)

	testUser := &entity.User{
		ID:        userID,
		FirstName: "Test",
		LastName:  "Test",
		Address:   nil,
	}

	jsonTestUser, _ := json.Marshal(testUser)

	mock.ExpectGet(testUser.ID).SetVal(string(jsonTestUser))

	user, err := UserRepository().GetUserByID(context.TODO(), userID)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), testUser, user)
}

// Get a user that does not exist
func (s *TestSuite) TestUserRepository_GetUserByID_Error() {
	userID := helper.UUID()
	db, mock := redismock.NewClientMock()
	storage.SetInstance(db)

	mock.ExpectGet(userID).RedisNil()

	_, err := UserRepository().GetUserByID(context.TODO(), userID)

	assert.NotNil(s.T(), err)
}

// Update a user that exist
func (s *TestSuite) TestUserRepo_UpdateUser_Success() {
	userID := helper.UUID()
	db, mock := redismock.NewClientMock()
	storage.SetInstance(db)

	testUser := &entity.User{
		ID:        userID,
		FirstName: "Test",
		LastName:  "Test",
		Address:   nil,
	}

	jsonTestUser, _ := json.Marshal(testUser)

	mock.ExpectSet(testUser.ID, string(jsonTestUser), redis.KeepTTL).SetVal(string(jsonTestUser))

	_, err := UserRepository().SaveUser(context.TODO(), testUser)

	assert.Nil(s.T(), err)
}

// Save a user that does not exist
func (s *TestSuite) TestUserRepo_SaveUser_Success() {
	db, _ := redismock.NewClientMock()
	storage.SetInstance(db)

	testUser := &entity.User{
		ID:        "",
		FirstName: "Test",
		LastName:  "Test",
		Address:   nil,
	}

	user, _ := UserRepository().SaveUser(context.TODO(), testUser)

	assert.NotEmpty(s.T(), user.ID)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
