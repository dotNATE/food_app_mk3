package handlers

import (
	"fmt"
	"main/models"
	utils "main/pkg"
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
		ctx.JSON(http.StatusBadRequest, utils.HTTPResponse{
			Success: false,
			Error:   "Invalid input",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Error hashing password: " + err.Error(),
		})
		return
	}

	user, err := handler.UserRepo.InsertUser(models.User{
		Email: input.Email,
		Name:  input.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Error creating user: " + err.Error(),
		})
		return
	}

	_, err = handler.AuthRepo.InsertAuthIdentity(models.AuthIdentity{
		UserId:   user.ID,
		Password: string(hashedPassword),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPResponse{
			Success: false,
			Error:   "Error creating authentication identity: " + err.Error(),
		})
		return
	}

	ctx.Header("Location", fmt.Sprintf("users/%d", user.ID))
	ctx.JSON(http.StatusCreated, utils.HTTPResponse{
		Success: true,
		Message: "User successfully created",
		Data:    user,
	})
}
