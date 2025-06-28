package handlers

import (
	"fmt"
	"main/dto"
	utils "main/pkg/utils"
	"main/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	app_url := os.Getenv("APP_URL")
	var input dto.RegisterRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input", nil))
		return
	}

	user, err := handler.UserService.RegisterNewUser(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Error registering new user: ", err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("%s/users/%d", app_url, user.ID)) // TODO implement this route!
	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("User successfully created", user))
}
