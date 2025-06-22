package main

import (
	"main/handlers"
	"main/pkg/middleware"
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
	authHandler := handlers.NewAuthHandler(authRepo, userRepo)

	router.POST("/auth/register", userHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	vendors := router.Group("/vendors")
	vendors.Use(middleware.JWTAuthMiddleware())

	vendors.GET("/", vendorHandler.GetVendors)
	vendors.POST("/", vendorHandler.AddNewVendor)
	vendors.GET("/:vendor_id", vendorHandler.GetVendorById)
	vendors.GET("/:vendor_id/ratings", ratingHandler.GetRatingsByVendorId)
	vendors.POST("/:vendor_id/ratings", ratingHandler.AddNewRating)
	vendors.GET("/:vendor_id/ratings/:rating_id", ratingHandler.GetRatingById)

	router.Run(":8080")
}
