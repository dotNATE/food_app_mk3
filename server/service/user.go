package service

import (
	"fmt"
	"main/dto"
	"main/models"
	"main/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo repository.UserRepositoryInterface
	AuthRepo repository.AuthRepositoryInterface
	DB       repository.DBInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface, authRepo repository.AuthRepositoryInterface, db repository.DBInterface) *UserService {
	return &UserService{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		DB:       db,
	}
}

func (service *UserService) RegisterNewUser(register_request *dto.RegisterRequest) (*dto.UserResponse, error) {
	user_exists, err := service.UserRepo.CheckUserExists(register_request.Email)
	if err != nil {
		return nil, err
	}
	if user_exists {
		return nil, fmt.Errorf("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register_request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user *models.User
	err = service.DB.WithTransaction(func(tx *gorm.DB) error {
		user, err = service.UserRepo.InsertUser(tx, models.User{
			Email: register_request.Email,
			Name:  register_request.Name,
		})
		if err != nil {
			return err
		}
		_, err = service.AuthRepo.InsertAuthIdentity(tx, models.AuthIdentity{
			UserId:   user.ID,
			Password: string(hashedPassword),
		})
		if err != nil {
			return err
		}

		return nil
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

func (service *UserService) GetUserById(user_id int64) (*dto.UserResponse, error) {
	user, err := service.UserRepo.GetUserById(user_id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
