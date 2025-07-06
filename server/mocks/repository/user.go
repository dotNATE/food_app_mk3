package repo_mocks

import (
	"main/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) InsertUser(tx *gorm.DB, user models.User) (*models.User, error) {
	args := mock.Called(tx, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (mock *MockUserRepository) CheckUserExists(email string) (bool, error) {
	args := mock.Called(email)
	return args.Bool(0), args.Error(1)
}

func (mock *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (mock *MockUserRepository) GetUserById(user_id int64) (*models.User, error) {
	args := mock.Called(user_id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (mock *MockUserRepository) GetDB() *gorm.DB {
	args := mock.Called()
	return args.Get(0).(*gorm.DB)
}
