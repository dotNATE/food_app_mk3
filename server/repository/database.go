package repository

import (
	"log"

	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:password@tcp(database:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate your schema
	if err := DB.AutoMigrate(&models.Vendor{}, &models.Rating{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return DB
}
