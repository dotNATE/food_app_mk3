package repository

import (
	"log"
	"os"

	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db_url := os.Getenv("DATABASE_URL")

	DB, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.Vendor{}, &models.Rating{}, &models.User{}, &models.AuthIdentity{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return DB
}
