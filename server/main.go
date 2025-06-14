package main

import (
	"main/handlers"
	"main/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitDB()
	defer repository.DB.Close()

	router := gin.Default()

	vendorHandler := handlers.NewVendorHandler(repository.NewVendorRepository(repository.DB))

	router.GET("/vendors", vendorHandler.GetVendors)
	router.POST("/vendors", vendorHandler.InsertNewVendor)

	router.Run(":8080")
}
