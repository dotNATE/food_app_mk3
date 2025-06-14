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
	ratingHandler := handlers.NewRatingHandler(repository.NewRatingRepository(repository.DB), repository.NewVendorRepository(repository.DB))

	router.GET("/vendors", vendorHandler.GetVendors)
	router.POST("/vendors", vendorHandler.AddNewVendor)
	router.GET("/vendors/:vendor_id", vendorHandler.GetVendorById)
	router.POST("/vendors/:vendor_id/ratings", ratingHandler.AddNewRating)
	// router.GET("/vendors/:vendor_id/ratings", ratingHandler.GetRatingsByVendorId) NOT YET IMPLEMENTED
	router.GET("/vendors/:vendor_id/ratings/:rating_id", ratingHandler.GetRatingById)

	router.Run(":8080")
}
