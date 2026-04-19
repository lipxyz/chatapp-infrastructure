package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("[WARNING] Failed to connect to Redis:", err)
	}
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
