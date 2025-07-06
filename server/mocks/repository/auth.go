package repo_mocks

import (
	"main/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockAuthRepository struct {
	mock.Mock
}

func (mock *MockAuthRepository) InsertAuthIdentity(tx *gorm.DB, auth models.AuthIdentity) (*models.AuthIdentity, error) {
	args := mock.Called(tx, auth)
	return args.Get(0).(*models.AuthIdentity), args.Error(1)
}

func (mock *MockAuthRepository) GetAuthByUserId(user_id int64) (*models.AuthIdentity, error) {
	args := mock.Called(user_id)
	return args.Get(0).(*models.AuthIdentity), args.Error(1)
}

func (mock *MockAuthRepository) GetDB() *gorm.DB {
	args := mock.Called()
	return args.Get(0).(*gorm.DB)
}
