package main

import (
	"log"
	"main/repository"
	"main/routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")

	db, err := repository.ConnectWithRetry(10, 2*time.Second)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	router := gin.Default()

	routes.RegisterRoutes(router, db)

	router.Run(":" + port)
}
