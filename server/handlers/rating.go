package handlers

import (
	"fmt"
	"main/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	RatingRepo *repository.RatingRepository
	VendorRepo *repository.VendorRepository
}

func NewRatingHandler(ratingRepo *repository.RatingRepository, vendorRepo *repository.VendorRepository) *RatingHandler {
	return &RatingHandler{RatingRepo: ratingRepo, VendorRepo: vendorRepo}
}

func (handler *RatingHandler) AddNewRating(ctx *gin.Context) {
	var new_rating repository.Rating
	vendor_id_param := ctx.Param("vendor_id")

	err := ctx.ShouldBindJSON(&new_rating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid input: " + err.Error(),
		})
		return
	}

	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid vendor_id, must be a number: " + err.Error(),
		})
		return
	}

	vendor_exists, err := handler.VendorRepo.CheckVendorExists(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Something went wrong checking vendor_id, please try again: " + err.Error(),
		})
		return
	}
	if !vendor_exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("No vendor found with id: %d", vendor_id),
		})
		return
	}

	new_rating.VendorId = vendor_id
	rating, err := handler.RatingRepo.InsertRating(new_rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to insert rating: " + err.Error(),
		})
		return
	}

	err = handler.VendorRepo.UpdateAverageRating(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update vendor average_rating: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Rating created successfully",
		"data":    rating,
	})
}
