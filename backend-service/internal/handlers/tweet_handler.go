package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sitter/internal/metrics"
	"sitter/internal/models"
	"time"
)

// Структура запроса на создание твита
type CreateTweetRequest struct {
	Text     string `json:"text" binding:"required"`
	UserID   string `json:"userid" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// Создание твита
// Создание твита
func CreateTweet(c *gin.Context) {
	start := time.Now() // Засекаем время начала создания твита

	var req CreateTweetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet := models.Tweet{
		ID:       uuid.New().String(),
		Text:     req.Text,
		UserID:   req.UserID,
		Username: req.Username,
		Likes:    0,
		Shares:   0,
		Slurmed:  0,
	}

	if err := tweet.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tweet"})
		return
	}

	// Увеличиваем метрику для создания твита
	metrics.TweetCreateTotal.Inc()
	metrics.TweetDuration.Observe(time.Since(start).Seconds())

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet created successfully", "tweet": tweet})
}

// Лайк твита
func LikeTweet(c *gin.Context) {
	start := time.Now() // Засекаем время начала лайка твита

	tweetID := c.Param("id")
	userID := c.PostForm("userID") // Получаем userID из формы или JSON-запроса

	// Получаем твит
	tweet, err := models.GetTweet(tweetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tweet"})
		return
	}

	// Лайк твита
	err = tweet.Like(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Увеличиваем метрику для лайка твита
	metrics.TweetLikeTotal.Inc()
	metrics.TweetDuration.Observe(time.Since(start).Seconds())

	c.JSON(http.StatusOK, gin.H{"message": "Tweet liked successfully"})
}

func SlurmTweet(c *gin.Context) {
	start := time.Now() // Засекаем время начала удаления твита

	tweetID := c.Param("id")
	userID := c.PostForm("userID") // Получаем userID из формы или JSON-запроса

	// Получаем твит
	tweet, err := models.GetTweet(tweetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tweet"})
		return
	}

	// Slurm твита
	err = tweet.Slurm(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Увеличиваем метрику для удаления твита
	metrics.TweetSlurmTotal.Inc()
	metrics.TweetDuration.Observe(time.Since(start).Seconds())

	c.JSON(http.StatusOK, gin.H{"message": "Tweet slurmed successfully"})
}

// Репост твита
func ShareTweet(c *gin.Context) {
	start := time.Now() // Засекаем время начала репоста

	tweetID := c.Param("id")

	tweet, err := models.GetTweet(tweetID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	tweet.Share()

	// Увеличиваем метрику для репоста твита
	metrics.TweetShareTotal.Inc()
	metrics.TweetDuration.Observe(time.Since(start).Seconds())

	c.JSON(http.StatusOK, gin.H{"message": "Tweet shared", "shares": tweet.Shares})
}

func GetAllTweets(c *gin.Context) {
	tweets, err := models.GetAllTweets()
	if err != nil {
		log.Println("Ошибка при получении твитов:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tweets"})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

func GetSortedTweets(c *gin.Context) {
	tweets, err := models.GetSortedTweetsByLikes()
	if err != nil {
		log.Println("Ошибка при получении отсортированных твитов:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tweets"})
		return
	}

	c.JSON(http.StatusOK, tweets)
}
func GetTweetByID(c *gin.Context) {
	tweetID := c.Param("id")

	tweet, err := models.GetTweet(tweetID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	c.JSON(http.StatusOK, tweet)
}
