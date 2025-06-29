package service

import (
	"main/dto"
	"main/models"
	"main/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
	AuthRepo *repository.AuthRepository
}

func NewUserService(userRepo *repository.UserRepository, authRepo *repository.AuthRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
		AuthRepo: authRepo,
	}
}

func (service *UserService) RegisterNewUser(register_request *dto.RegisterRequest) (*dto.UserResponse, error) {
	user, err := service.UserRepo.InsertUser(models.User{
		Email: register_request.Email,
		Name:  register_request.Name,
	})
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register_request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = service.AuthRepo.InsertAuthIdentity(models.AuthIdentity{
		UserId:   user.ID,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (service *UserService) GetUserByEmail(email string) (*dto.UserResponse, error) {
	user, err := service.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
