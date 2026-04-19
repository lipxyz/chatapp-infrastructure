package database

import (
	"log"
	"sitter/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("[WARNING] Migration failed:", err)
	} else {
		log.Println("[SUCCESS] Database migrated successfully!")
	}
}
