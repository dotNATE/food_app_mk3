package handlers

import (
	"fmt"
	utils "main/pkg"
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
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid vendor_id, must be a number: " + err.Error(),
		})
		return
	}

	vendor_exists, err := handler.VendorRepo.CheckVendorExists(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Something went wrong checking vendor_id, please try again: " + err.Error(),
		})
		return
	}
	if !vendor_exists {
		ctx.JSON(http.StatusNotFound, utils.HTTPResponse{
			Success: false,
			Error:   fmt.Sprintf("No vendor found with id: %d", vendor_id),
		})
		return
	}

	new_rating.VendorId = vendor_id
	rating, err := handler.RatingRepo.InsertRating(new_rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Failed to insert rating: " + err.Error(),
		})
		return
	}

	err = handler.VendorRepo.UpdateAverageRating(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Failed to update vendor average_rating: " + err.Error(),
		})
		return
	}

	ctx.Header("Location", fmt.Sprintf("vendors/%d/ratings/%d", rating.VendorId, rating.ID))
	ctx.JSON(http.StatusCreated, utils.HTTPResponse{
		Success: true,
		Message: "Rating created successfully",
		Data:    rating,
	})
}

func (handler *RatingHandler) GetRatingById(ctx *gin.Context) {
	rating_id_param := ctx.Param("rating_id")
	rating_id, err := strconv.ParseInt(rating_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid rating id, must be a number: " + err.Error(),
		})
		return
	}

	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid vendor id, must be a number: " + err.Error(),
		})
		return
	}

	rating, err := handler.RatingRepo.GetRatingById(rating_id, vendor_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.HTTPResponse{
		Success: true,
		Message: "Successfully fetched rating",
		Data:    rating,
	})
}

func (handler *RatingHandler) GetRatingsByVendorId(ctx *gin.Context) {
	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid vendor_id",
		})
		return
	}

	exists, err := handler.VendorRepo.CheckVendorExists(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Something went wrong checking vendor_id, please try again: " + err.Error(),
		})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, utils.HTTPResponse{
			Success: false,
			Error:   "Vendor not found",
		})
		return
	}

	ratings, err := handler.RatingRepo.GetRatingsByVendorId(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Failed to fetch ratings" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.HTTPResponse{
		Success: true,
		Message: "Successfully fetched Ratings",
		Data:    ratings,
	})
}
