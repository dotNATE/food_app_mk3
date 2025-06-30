package handlers

import (
	"fmt"
	"main/dto"
	utils "main/pkg/utils"
	"main/service"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	VendorService *service.VendorService
	RatingService *service.RatingService
}

func NewRatingHandler(ratingService *service.RatingService, vendorService *service.VendorService) *RatingHandler {
	return &RatingHandler{RatingService: ratingService, VendorService: vendorService}
}

func (handler *RatingHandler) AddNewRating(ctx *gin.Context) {
	app_url := os.Getenv("APP_URL")
	var new_rating *dto.AddRatingRequest
	vendor_id_param := ctx.Param("vendor_id")

	err := ctx.ShouldBindJSON(&new_rating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input: ", err))
		return
	}

	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid vendor_id, must be a number: ", err))
		return
	}

	vendor_exists, err := handler.VendorService.CheckVendorExists(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Something went wrong checking vendor_id, please try again: ", err))
		return
	}
	if !vendor_exists {
		ctx.JSON(http.StatusNotFound, utils.CreateErrorHTTPResponse(fmt.Sprintf("No vendor found with id: %d", vendor_id), nil))
		return
	}

	rating, err := handler.RatingService.CreateNewRating(new_rating, vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Failed to insert rating: ", err))
		return
	}

	err = handler.VendorService.UpdateAverageRating(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Failed to update vendor average_rating: ", err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("%s/vendors/%d/ratings/%d", app_url, rating.VendorId, rating.ID))
	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("Rating created successfully", rating))
}

func (handler *RatingHandler) GetRatingById(ctx *gin.Context) {
	rating_id_param := ctx.Param("rating_id")
	rating_id, err := strconv.ParseInt(rating_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid rating id, must be a number: ", err))
		return
	}

	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid vendor id, must be a number: ", err))
		return
	}

	rating, err := handler.RatingService.GetRatingById(rating_id, vendor_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Failed to fetch rating: ", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched rating", rating))
}

func (handler *RatingHandler) GetRatingsByVendorId(ctx *gin.Context) {
	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid vendor_id", nil))
		return
	}

	exists, err := handler.VendorService.CheckVendorExists(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Something went wrong checking vendor_id, please try again: ", err))
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, utils.CreateErrorHTTPResponse("Vendor not found", nil))
		return
	}

	ratings, err := handler.RatingService.GetRatingsByVendorId(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Failed to fetch ratings", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched Ratings", ratings))
}
