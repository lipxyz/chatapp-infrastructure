package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[WARNING] No .env file found")
	} else {
		log.Println("[SUCCESS] .env file loaded successfully")
	}

	//fmt.Println("POSTGRES_HOST:", os.Getenv("POSTGRES_HOST"))
	//fmt.Println("POSTGRES_USER:", os.Getenv("POSTGRES_USER"))
	//fmt.Println("POSTGRES_PASSWORD:", os.Getenv("POSTGRES_PASSWORD"))
	//fmt.Println("POSTGRES_DB:", os.Getenv("POSTGRES_DB"))
	//fmt.Println("POSTGRES_PORT:", os.Getenv("POSTGRES_PORT"))

	if os.Getenv("POSTGRES_HOST") == "" || os.Getenv("POSTGRES_PORT") == "" {
		log.Fatal("[WARNING] Missing required environment variables!")
	}
	var dsn string
	if os.Getenv("POSTGRES_PASSWORD") == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PORT"),
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PORT"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("[WARNING] Failed to connect to PostgreSQL:", err)
	}

	DB = db
}
func GetPostgresConnString() string {
	// Выводим переменные окружения для проверки
	//fmt.Println("POSTGRES_HOST:", os.Getenv("POSTGRES_HOST"))
	//fmt.Println("POSTGRES_USER:", os.Getenv("POSTGRES_USER"))
	//fmt.Println("POSTGRES_PASSWORD:", os.Getenv("POSTGRES_PASSWORD"))
	//fmt.Println("POSTGRES_DB:", os.Getenv("POSTGRES_DB"))
	//fmt.Println("POSTGRES_PORT:", os.Getenv("POSTGRES_PORT"))

	// Проверка наличия необходимых переменных окружения
	if os.Getenv("POSTGRES_HOST") == "" || os.Getenv("POSTGRES_PORT") == "" {
		return ""
	}

	// Формируем строку подключения для PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	return dsn
}
