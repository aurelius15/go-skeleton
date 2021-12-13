package storage

import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func Instance() *redis.Client {
	return redisClient
}

func SetInstance(client *redis.Client) {
	redisClient = client
}
