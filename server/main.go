package main

import (
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVendors(ctx *gin.Context) {
	vendors, err := repository.GetAllVendors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "There was a problem fetching the vendors: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully fetched all vendors",
		"vendors": vendors,
	})
}

func insertNewVendor(ctx *gin.Context) {
	var newVendor repository.Vendor

	err := ctx.ShouldBindJSON(&newVendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid input: " + err.Error(),
		})
		return
	}

	id, err := repository.InsertVendor(newVendor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to insert vendor: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":   true,
		"message":   "Vendor created successfully",
		"vendor_id": id,
	})
}

func main() {
	repository.InitDB()
	defer repository.DB.Close()

	router := gin.Default()

	router.GET("/vendors", getVendors)
	router.POST("/vendors", insertNewVendor)

	router.Run(":8080")
}
