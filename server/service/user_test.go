package service_test

import (
	"errors"
	"main/dto"
	repoMocks "main/mocks/repository"
	"main/models"
	"main/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterNewUser_Success(t *testing.T) {
	mockUserRepo := new(repoMocks.MockUserRepository)
	mockAuthRepo := new(repoMocks.MockAuthRepository)
	mockGormWrapper := new(repoMocks.MockGormDBWrapper)

	user := models.User{
		ID:    1,
		Email: "test@example.com",
		Name:  "John Doe",
	}
	auth_identity := models.AuthIdentity{
		UserId:   1,
		Password: "hashed",
	}

	mockUserRepo.On("CheckUserExists", "test@example.com").Return(false, nil)
	mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(&user, nil)
	mockAuthRepo.On("InsertAuthIdentity", mock.Anything, mock.Anything).Return(&auth_identity, nil)

	service := &service.UserService{
		DB:       mockGormWrapper,
		UserRepo: mockUserRepo,
		AuthRepo: mockAuthRepo,
	}

	result, err := service.RegisterNewUser(dto.NewRegisterRequest(
		"John Doe",
		"test@example.com",
		"secret123",
	))

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "test@example.com", result.Email)
	assert.Equal(t, "John Doe", result.Name)

	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
}

func TestGetUserByEmail_Success(t *testing.T) {
	mockUserRepo := new(repoMocks.MockUserRepository)
	userModel := &models.User{
		ID:    42,
		Email: "id@example.com",
		Name:  "ID User",
	}
	expectedUser := &dto.UserResponse{
		ID:    42,
		Email: "id@example.com",
		Name:  "ID User",
	}

	mockUserRepo.On("GetUserByEmail", "test@example.com").Return(userModel, nil)

	userService := &service.UserService{
		UserRepo: mockUserRepo,
	}

	user, err := userService.GetUserByEmail("test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByEmail_Error(t *testing.T) {
	mockUserRepo := new(repoMocks.MockUserRepository)

	mockUserRepo.On("GetUserByEmail", "notfound@example.com").Return((*models.User)(nil), errors.New("user not found"))

	userService := &service.UserService{
		UserRepo: mockUserRepo,
	}

	user, err := userService.GetUserByEmail("notfound@example.com")

	assert.Nil(t, user)
	assert.EqualError(t, err, "user not found")
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserById_Success(t *testing.T) {
	mockUserRepo := new(repoMocks.MockUserRepository)
	userModel := &models.User{
		ID:    42,
		Email: "id@example.com",
		Name:  "ID User",
	}
	expectedUser := &dto.UserResponse{
		ID:    42,
		Email: "id@example.com",
		Name:  "ID User",
	}

	mockUserRepo.On("GetUserById", int64(42)).Return(userModel, nil)

	userService := &service.UserService{
		UserRepo: mockUserRepo,
	}

	user, err := userService.GetUserById(42)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserById_Error(t *testing.T) {
	mockUserRepo := new(repoMocks.MockUserRepository)
	mockUserRepo.On("GetUserById", int64(999)).Return((*models.User)(nil), errors.New("user not found"))

	userService := &service.UserService{
		UserRepo: mockUserRepo,
	}

	user, err := userService.GetUserById(999)

	assert.Nil(t, user)
	assert.EqualError(t, err, "user not found")
	mockUserRepo.AssertExpectations(t)
}
