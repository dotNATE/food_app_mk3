package handlers

import (
	"fmt"
	"main/models"
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
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("There was a problem fetching the vendors: ", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched all vendors", vendors))
}

func (handler *VendorHandler) AddNewVendor(ctx *gin.Context) {
	var new_vendor models.Vendor

	err := ctx.ShouldBindJSON(&new_vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input: ", err))
		return
	}

	vendor, err := handler.VendorRepo.InsertVendor(new_vendor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Failed to insert vendor: ", err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("vendors/%d", vendor.ID))
	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("Vendor created successfully", vendor))
}

func (handler *VendorHandler) GetVendorById(ctx *gin.Context) {
	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid vendor id, must be a number: ", err))
		return
	}

	vendor, err := handler.VendorRepo.GetVendorById(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Failed fetching vendor: ", err))
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched vendor", vendor))
}
