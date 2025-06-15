package handlers

import (
	utils "main/pkg"
	"main/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthRepo *repository.AuthRepository
	UserRepo *repository.UserRepository
}

func NewAuthHandler(authRepo *repository.AuthRepository, userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo, UserRepo: userRepo}
}

var JwtSecret = []byte("my-secret") // TODO move this to .env

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	user, err := handler.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Something went wrong checking user, please try again: " + err.Error(),
		})
		return
	}

	auth_identity, err := handler.AuthRepo.GetAuthByUserId(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Something went wrong checking auth credentials, please try again: " + err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth_identity.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.HTTPResponse{
			Success: false,
			Message: "Invalid credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret) // TODO NATE obviously move this to a .env
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Message: "Token creation failed: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.HTTPResponse{
		Success: true,
		Data:    tokenString,
	})
}
