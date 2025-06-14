package handlers

import (
	utils "main/pkg"
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
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "There was a problem fetching the vendors: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.HTTPResponse{
		Success: true,
		Message: "Successfully fetched all vendors",
		Data:    vendors,
	})
}

func (handler *VendorHandler) AddNewVendor(ctx *gin.Context) {
	var new_vendor repository.Vendor

	err := ctx.ShouldBindJSON(&new_vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	vendor, err := handler.VendorRepo.InsertVendor(new_vendor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Failed to insert vendor: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.HTTPResponse{
		Success: true,
		Message: "Vendor created successfully",
		Data:    vendor,
	})
}
