package repository

import (
	"context"
	"encoding/json"

	"github.com/aurelius15/go-skeleton/internal/entity"
	"github.com/aurelius15/go-skeleton/internal/helper"
	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/go-redis/redis/v8"
)

var userRepositoryInstance *UserRepo

type UserRepo struct {
	client *redis.Client
}

func UserRepository() *UserRepo {
	if userRepositoryInstance == nil {
		userRepositoryInstance = &UserRepo{client: storage.Instance()}
	}

	return userRepositoryInstance
}

func (r *UserRepo) NumberOfUsers(ctx context.Context) (n int64, err error) {
	n, err = r.client.DBSize(ctx).Result()
	return
}

func (r *UserRepo) GetUserByID(ctx context.Context, userID string) (u *entity.User, err error) {
	user, err := r.client.Get(ctx, userID).Result()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(user), &u)

	return
}

func (r *UserRepo) SaveUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.ID == "" {
		user.ID = helper.UUID()
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return user, err
	}

	_, err = r.client.Set(ctx, user.ID, string(jsonUser), redis.KeepTTL).Result()

	if err != nil {
		return user, err
	}

	return user, nil
}
