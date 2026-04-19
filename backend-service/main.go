package main

import (
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sitter/internal/metrics"
	"strconv"
	"strings"

	"sitter/internal/database"
	"sitter/internal/database/redisclient"
	"sitter/router"
)

func main() {
	// Загружаем переменные окружения из .env (если файл есть)
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используем переменные окружения")
	}

	// Инициализация PostgreSQL
	database.InitPostgres()

	// Инициализация Redis
	database.InitRedis()

	// Миграции
	allowMigrate := strings.EqualFold(os.Getenv("MIGRATE_IN_CODE"), "true")

	if allowMigrate {
		database.Migrate()
	} else {
		log.Println("[WARNING] MIGRATE_IN_CODE WAS SKIPPED. WATCH ENV -> MIGRATE_IN_CODE")
	}

	metrics.SimulatedMetrics()

	// Получаем настройки Redis из env
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	// Инициализация Redis клиента
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	redisclient.InitRedisClient(client)

	// Подключаем основное API
	r := router.SetupRouter()

	// Разрешаем CORS для всех методов и всех источников
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},                                       // Разрешаем доступ только с localhost:3000
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем необходимые методы
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем необходимые заголовки
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	r.Run(":8080")
}
