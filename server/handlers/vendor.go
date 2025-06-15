package handlers

import (
	"fmt"
	utils "main/pkg"
	"main/repository"
	"net/http"
	"strconv"

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

	ctx.Header("Location", fmt.Sprintf("vendors/%d", vendor.ID))
	ctx.JSON(http.StatusCreated, utils.HTTPResponse{
		Success: true,
		Message: "Vendor created successfully",
		Data:    vendor,
	})
}

func (handler *VendorHandler) GetVendorById(ctx *gin.Context) {
	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid vendor id, must be a number: " + err.Error(),
		})
		return
	}

	vendor, err := handler.VendorRepo.GetVendorById(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Failed fetching vendor: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.HTTPResponse{
		Success: true,
		Message: "Successfully fetched vendor",
		Data:    vendor,
	})
}
