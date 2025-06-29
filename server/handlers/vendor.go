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

type VendorHandler struct {
	VendorService *service.VendorService
}

func NewVendorHandler(vendorService *service.VendorService) *VendorHandler {
	return &VendorHandler{
		VendorService: vendorService,
	}
}

func (handler *VendorHandler) GetVendors(ctx *gin.Context) {
	vendors, err := handler.VendorService.GetAllVendors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("There was a problem fetching the vendors: ", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched all vendors", vendors))
}

func (handler *VendorHandler) AddNewVendor(ctx *gin.Context) {
	app_url := os.Getenv("APP_URL")

	var new_vendor *dto.AddVendorRequest
	err := ctx.ShouldBindJSON(&new_vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input: ", err))
		return
	}

	user_id, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Issue getting user_id from context: ", err))
		return
	}

	vendor, err := handler.VendorService.CreateNewVendor(new_vendor, user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Failed to insert vendor: ", err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("%s/vendors/%d", app_url, vendor.ID))
	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("Vendor created successfully", vendor))
}

func (handler *VendorHandler) GetVendorById(ctx *gin.Context) {
	vendor_id_param := ctx.Param("vendor_id")
	vendor_id, err := strconv.ParseInt(vendor_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid vendor id, must be a number: ", err))
		return
	}

	vendor, err := handler.VendorService.GetVendorById(vendor_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Failed fetching vendor: ", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Successfully fetched vendor", vendor))
}
