package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) InsertUser(user models.User) (*models.User, error) {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) CheckUserExists(email string) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return count > 0, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUserById(user_id int64) (*models.User, error) {
	var user models.User

	err := repo.DB.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
