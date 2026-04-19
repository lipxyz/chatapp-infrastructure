package router

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sitter/internal/database"
	"sitter/internal/database/redisclient"
	"sitter/internal/handlers"
	"sitter/internal/metrics"
)

// Функция для Liveness Probe (просто проверка, что сервер жив)
func livenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Функция для Readiness Probe (проверяет подключение к Redis)
func readinessCheck(c *gin.Context) {
	ctx := context.Background()

	// Проверка Redis
	_, err := redisclient.GetRedisClient().Ping(ctx).Result()
	if err != nil {
		log.Println("Redis недоступен:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Redis unavailable"})
		return
	}

	// Проверка PostgreSQL
	conn, err := pgx.Connect(ctx, database.GetPostgresConnString())
	if err != nil {
		log.Println("PostgreSQL недоступен:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "PostgreSQL unavailable"})
		return
	}
	defer conn.Close(ctx)

	// Простая проверка доступности (например, выполнение запроса)
	err = conn.Ping(ctx)
	if err != nil {
		log.Println("PostgreSQL недоступен:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "PostgreSQL unavailable"})
		return
	}

	// Если оба сервиса доступны
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},                                       // Разрешаем доступ только с localhost:3000
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем необходимые методы
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем необходимые заголовки
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.Use(metrics.MetricsMiddleware())
	api := r.Group("/api")
	{
		api.Use(func(c *gin.Context) {
			c.Set("db", database.DB) // Добавляем объект DB в контекст
			c.Next()
		})
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)
		api.GET("/users", handlers.GetAllUsersHandler)

		// Твиты
		api.GET("/tweets", handlers.GetAllTweets)
		api.POST("/tweets", handlers.CreateTweet)
		api.GET("/tweets/:id", handlers.GetTweetByID)
		api.POST("/tweets/:id/like", handlers.LikeTweet)
		api.POST("/tweets/:id/share", handlers.ShareTweet)
		api.POST("/tweets/:id/slurm", handlers.SlurmTweet)
		api.GET("/tweets/sorted", handlers.GetSortedTweets)

		// Добавляем маршруты для Liveness, Readiness и метрик
		r.GET("/healthz", livenessCheck)
		r.GET("/readyz", readinessCheck)
		r.GET("/metrics", gin.WrapH(promhttp.Handler())) // Метрики Prometheus
	}

	return r
}
