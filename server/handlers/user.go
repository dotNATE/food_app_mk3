package handlers

import (
	"fmt"
	"main/models"
	utils "main/pkg/utils"
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
	AuthRepo *repository.AuthRepository
}

func NewUserHandler(userRepo *repository.UserRepository, authRepo *repository.AuthRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
		AuthRepo: authRepo,
	}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorHTTPResponse("Invalid input", nil))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Error hashing password: ", err))
		return
	}

	user, err := handler.UserRepo.InsertUser(models.User{
		Email: input.Email,
		Name:  input.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Error creating user: ", err))
		return
	}

	_, err = handler.AuthRepo.InsertAuthIdentity(models.AuthIdentity{
		UserId:   user.ID,
		Password: string(hashedPassword),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorHTTPResponse("Error creating authentication identity: ", err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("users/%d", user.ID))
	ctx.JSON(http.StatusCreated, utils.CreateSuccessfulHTTPResponse("User successfully created", user))
}
