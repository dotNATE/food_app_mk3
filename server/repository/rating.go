package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type RatingRepository struct {
	DB *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{DB: db}
}

func (repo *RatingRepository) InsertRating(rating models.Rating) (*models.Rating, error) {
	result := repo.DB.Create(&rating)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to insert rating: %w", result.Error)
	}

	return &rating, nil
}

func (repo *RatingRepository) GetRatingById(rating_id int64, vendor_id int64) (*models.Rating, error) {
	var rating models.Rating

	err := repo.DB.Where("id = ? AND vendor_id = ?", rating_id, vendor_id).First(&rating).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rating: %w", err)
	}

	return &rating, nil
}

func (repo *RatingRepository) GetRatingsByVendorId(vendor_id int64) (*[]models.Rating, error) {
	var ratings []models.Rating

	err := repo.DB.Where("vendor_id = ?", vendor_id).Find(&ratings).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ratings: %w", err)
	}

	return &ratings, nil
}
