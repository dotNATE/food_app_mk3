package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type RatingRepository interface {
	InsertRating(rating models.Rating) (*models.Rating, error)
	GetRatingById(rating_id int64, vendor_id int64) (*models.Rating, error)
	GetRatingsByVendorId(vendor_id int64) ([]*models.Rating, error)
	GetDB() *gorm.DB
}

type GormRatingRepository struct {
	DB *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return GormRatingRepository{DB: db}
}

func (repo GormRatingRepository) InsertRating(rating models.Rating) (*models.Rating, error) {
	err := repo.DB.Create(&rating).Error
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (repo GormRatingRepository) GetRatingById(rating_id int64, vendor_id int64) (*models.Rating, error) {
	var rating models.Rating

	err := repo.DB.Where("id = ? AND vendor_id = ?", rating_id, vendor_id).First(&rating).Error
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (repo GormRatingRepository) GetRatingsByVendorId(vendor_id int64) ([]*models.Rating, error) {
	var ratings []*models.Rating

	err := repo.DB.Where("vendor_id = ?", vendor_id).Find(&ratings).Error
	if err != nil {
		return nil, err
	}

	return ratings, nil
}

func (repo GormRatingRepository) GetDB() *gorm.DB {
	return repo.DB
}
