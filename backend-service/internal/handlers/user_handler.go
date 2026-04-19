package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"sitter/internal/auth"
	"sitter/internal/database"
	"sitter/internal/metrics"
	"sitter/internal/models"
	"time"
)

// Структура запроса для логина
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=1"`
}

// Регистрация пользователя
func RegisterUser(c *gin.Context) {
	start := time.Now() // Засекаем время начала регистрации

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Пароль будет захеширован автоматически
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Регистрируем метрику регистрации
	metrics.RegisterTotal.Inc()
	metrics.RegisterDuration.Observe(time.Since(start).Seconds())

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Логин пользователя и генерация JWT
func LoginUser(c *gin.Context) {
	start := time.Now() // Засекаем время начала логина

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		// Неудачный логин
		metrics.LoginFailTotal.Inc()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Сравниваем хешированный пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		// Неудачный логин
		metrics.LoginFailTotal.Inc()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Успешный логин
	metrics.LoginSuccessTotal.Inc()
	metrics.LoginDuration.Observe(time.Since(start).Seconds())

	// Генерация JWT
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Возвращаем токен
	c.JSON(http.StatusOK, gin.H{"token": token, "userId": user.ID, "username": user.Username, "email": user.Email})
}

func GetAllUsersHandler(c *gin.Context) {
	// Получаем объект DB из контекста
	dbInterface, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	// Приводим его к типу *gorm.DB
	db, ok := dbInterface.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert database type"})
		return
	}

	// Получаем пользователей из базы
	users, err := models.GetAllUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	// Возвращаем пользователей
	c.JSON(http.StatusOK, users)
}
