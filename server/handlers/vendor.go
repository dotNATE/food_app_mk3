package handlers

import (
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	VendorRepo *repository.VendorRepository
}

func NewVendorHandler(vendorRepo *repository.VendorRepository) *VendorHandler {
	return &VendorHandler{VendorRepo: vendorRepo}
}

func (handler *VendorHandler) GetVendors(ctx *gin.Context) {
	vendors, err := handler.VendorRepo.GetAllVendors()
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

func (handler *VendorHandler) InsertNewVendor(ctx *gin.Context) {
	var new_vendor repository.Vendor

	err := ctx.ShouldBindJSON(&new_vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid input: " + err.Error(),
		})
		return
	}

	id, err := handler.VendorRepo.InsertVendor(new_vendor)
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
