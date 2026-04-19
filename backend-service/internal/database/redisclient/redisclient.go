// internal/database/redisclient.go
package redisclient

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// Глобальная переменная для Redis
var RedisClient *redis.Client

// Инициализация клиента Redis
func InitRedisClient(client *redis.Client) {
	RedisClient = client
}

// Получение Redis клиента
func GetRedisClient() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis client is not initialized")
	}
	return RedisClient
}

// Получение контекста
func GetContext() context.Context {
	return context.Background()
}
