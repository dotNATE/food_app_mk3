package main

import (
	"main/handlers"
	"main/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.InitDB()

	router := gin.Default()

	vendorRepo := repository.NewVendorRepository(db)
	ratingRepo := repository.NewRatingRepository(db)
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)

	vendorHandler := handlers.NewVendorHandler(vendorRepo)
	ratingHandler := handlers.NewRatingHandler(ratingRepo, vendorRepo)
	userHandler := handlers.NewUserHandler(userRepo, authRepo)

	router.POST("/users", userHandler.Register)

	router.GET("/vendors", vendorHandler.GetVendors)
	router.POST("/vendors", vendorHandler.AddNewVendor)
	router.GET("/vendors/:vendor_id", vendorHandler.GetVendorById)
	router.GET("/vendors/:vendor_id/ratings", ratingHandler.GetRatingsByVendorId)
	router.POST("/vendors/:vendor_id/ratings", ratingHandler.AddNewRating)
	router.GET("/vendors/:vendor_id/ratings/:rating_id", ratingHandler.GetRatingById)

	router.Run(":8080")
}
