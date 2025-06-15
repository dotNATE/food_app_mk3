package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (repo *AuthRepository) InsertAuthIdentity(auth models.AuthIdentity) (*models.AuthIdentity, error) {
	err := repo.DB.Create(&auth).Error
	if err != nil {
		return nil, fmt.Errorf("failed to insert auth identity: %w", err)
	}

	return &auth, nil
}

func (repo *AuthRepository) GetAuthByUserId(user_id int64) (*models.AuthIdentity, error) {
	var auth_identity models.AuthIdentity

	err := repo.DB.First(&auth_identity).Where("user_id = ?", user_id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch auth identity: %w", err)
	}

	return &auth_identity, nil
}
