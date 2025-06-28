package handlers

import (
	utils "main/pkg/utils"
	"main/repository"
	"net/http"
	"os"
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

func (handler *AuthHandler) Login(ctx *gin.Context) {
	jwt_secret := []byte(os.Getenv("JWT_SECRET"))

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input: ", err))
		return
	}

	user, err := handler.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Something went wrong checking user, please try again: ", err))
		return
	}

	auth_identity, err := handler.AuthRepo.GetAuthByUserId(user.ID)
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
