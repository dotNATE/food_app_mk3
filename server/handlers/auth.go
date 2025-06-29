package handlers

import (
	"main/dto"
	utils "main/pkg/utils"
	"main/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserService *service.UserService
	AuthService *service.AuthService
}

func NewAuthHandler(userService *service.UserService, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{UserService: userService, AuthService: authService}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	jwt_secret := []byte(os.Getenv("JWT_SECRET"))

	var input *dto.LoginRequest
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input: ", err))
		return
	}

	user, err := handler.UserService.GetUserByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Something went wrong checking user, please try again: ", err))
		return
	}

	auth_identity, err := handler.AuthService.GetAuthIdentityByUserId(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Something went wrong checking auth credentials, please try again: ", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth_identity.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorHTTPResponse("Invalid credentials", nil))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(jwt_secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Token creation failed: ", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.CreateSuccessfulHTTPResponse("Login successful", tokenString))
}
