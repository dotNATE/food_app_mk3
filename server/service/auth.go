package service

import (
	"main/dto"
	"main/repository"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(userRepo *repository.UserRepository, authRepo *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepo: authRepo,
	}
}

func (service *AuthService) GetAuthIdentityByUserId(user_id int64) (*dto.AuthIdentity, error) {
	auth_identity, err := service.AuthRepo.GetAuthByUserId(user_id)
	if err != nil {
		return nil, err
	}

	return &dto.AuthIdentity{
		ID:       auth_identity.ID,
		UserId:   auth_identity.UserId,
		Password: auth_identity.Password,
	}, nil
}
