package routes

import (
	"main/handlers"
	"main/pkg/middleware"
	"main/repository"
	"main/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	vendorRepo := repository.NewVendorRepository(db)
	ratingRepo := repository.NewRatingRepository(db)
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)

	userService := service.NewUserService(userRepo, authRepo)
	authService := service.NewAuthService(userRepo, authRepo)
	vendorService := service.NewVendorService(vendorRepo)
	ratingService := service.NewRatingService(ratingRepo)

	vendorHandler := handlers.NewVendorHandler(vendorService)
	ratingHandler := handlers.NewRatingHandler(ratingService, vendorService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService, authService)

	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	vendors := router.Group("/vendors")
	vendors.Use(middleware.JWTAuthMiddleware())
	{
		vendors.GET("/", vendorHandler.GetVendors)
		vendors.POST("/", vendorHandler.AddNewVendor)
		vendors.GET("/:vendor_id", vendorHandler.GetVendorById)
		vendors.GET("/:vendor_id/ratings", ratingHandler.GetRatingsByVendorId)
		vendors.POST("/:vendor_id/ratings", ratingHandler.AddNewRating)
		vendors.GET("/:vendor_id/ratings/:rating_id", ratingHandler.GetRatingById)
	}
}
