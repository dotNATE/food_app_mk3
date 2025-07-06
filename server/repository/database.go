package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInterface interface {
	WithTransaction(func(tx *gorm.DB) error) error
	GetDB() *gorm.DB
}

type GormDB struct {
	DB *gorm.DB
}

func ConnectWithRetry(maxRetries int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	db_url := os.Getenv("DATABASE_URL")

	for attempts := 1; attempts <= maxRetries; attempts++ {
		db, err = gorm.Open(mysql.Open(db_url), &gorm.Config{})
		if err == nil {
			// Optional: run a simple query to ensure the connection is healthy
			sqlDB, errPing := db.DB()
			if errPing == nil {
				errPing = sqlDB.Ping()
			}
			if errPing == nil {
				log.Println("Connected to the database with GORM.")
				return db, nil
			}
			log.Printf("Attempt %d: Ping failed: %v", attempts, errPing)
		} else {
			log.Printf("Attempt %d: GORM open failed: %v", attempts, err)
		}

		time.Sleep(delay)
	}

	err = db.AutoMigrate(&models.Vendor{}, &models.Rating{}, &models.User{}, &models.AuthIdentity{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts: %v", maxRetries, err)
}

func (g GormDB) WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := g.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.DB
}
