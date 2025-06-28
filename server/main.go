package main

import (
	"log"
	"main/repository"
	"main/routes"
	"os"

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

	db := repository.InitDB()
	router := gin.Default()

	routes.RegisterRoutes(router, db)

	router.Run(":" + port)
}
