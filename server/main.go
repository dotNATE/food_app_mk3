package main

import (
	"fmt"
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVendors(ctx *gin.Context) {
	vendors, err := repository.GetAllVendors()
	if err != nil {
		fmt.Printf("failed to query vendors: %+v", err)
	}

	ctx.JSON(http.StatusOK, vendors)
}

func main() {
	repository.InitDB()
	defer repository.DB.Close()

	router := gin.Default()

	router.GET("/vendors", getVendors)

	router.Run(":8080")
}
