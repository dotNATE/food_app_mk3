package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	InsertAuthIdentity(tx *gorm.DB, auth models.AuthIdentity) (*models.AuthIdentity, error)
	GetAuthByUserId(user_id int64) (*models.AuthIdentity, error)
	GetDB() *gorm.DB
}

type GormAuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepositoryInterface {
	return GormAuthRepository{DB: db}
}

func (repo GormAuthRepository) InsertAuthIdentity(tx *gorm.DB, auth models.AuthIdentity) (*models.AuthIdentity, error) {
	err := tx.Create(&auth).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert auth identity: %w", err)
	}

	return &auth, nil
}

func (repo GormAuthRepository) GetAuthByUserId(user_id int64) (*models.AuthIdentity, error) {
	var auth_identity models.AuthIdentity

	err := repo.DB.First(&auth_identity).Where("user_id = ?", user_id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch auth identity: %w", err)
	}

	return &auth_identity, nil
}

func (repo GormAuthRepository) GetDB() *gorm.DB {
	return repo.DB
}
