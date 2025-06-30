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

func (handler *UserHandler) GetUserById(ctx *gin.Context) {
	user_id_param := ctx.Param("user_id")
	user_id, err := strconv.ParseInt(user_id_param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid user id, must be a number: ", err))
		return
	}

	user, err := handler.UserService.GetUserById(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Error fetching user: ", err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("User successfully retrieved", user))
}
