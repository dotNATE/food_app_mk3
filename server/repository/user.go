package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	InsertUser(tx *gorm.DB, user models.User) (*models.User, error)
	CheckUserExists(email string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(user_id int64) (*models.User, error)
	GetDB() *gorm.DB
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return GormUserRepository{DB: db}
}

func (repo GormUserRepository) InsertUser(tx *gorm.DB, user models.User) (*models.User, error) {
	err := tx.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo GormUserRepository) CheckUserExists(email string) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return count > 0, nil
}

func (repo GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo GormUserRepository) GetUserById(user_id int64) (*models.User, error) {
	var user models.User

	err := repo.DB.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo GormUserRepository) GetDB() *gorm.DB {
	return repo.DB
}
