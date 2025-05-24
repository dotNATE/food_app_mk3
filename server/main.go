package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type vendor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var vendors = []vendor{
	{Id: "ecfe5072-4a00-46c0-83a2-c54fabca4ce5", Name: "Three Brothers Burgers"},
	{Id: "d19ddb3f-2733-42f2-b140-e2817fe624ab", Name: "Chido Wey"},
}

func getVendors(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, vendors)
}

func main() {
	router := gin.Default()

	router.GET("/vendors", getVendors)

	router.Run("localhost:8080")
}
