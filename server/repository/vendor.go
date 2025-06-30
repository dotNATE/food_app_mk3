package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type VendorRepository struct {
	DB *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{DB: db}
}

func (repo *VendorRepository) GetAllVendors() ([]*models.Vendor, error) {
	var vendors []*models.Vendor

	err := repo.DB.Find(&vendors).Error
	if err != nil {
		return nil, err
	}

	return vendors, nil
}

func (repo *VendorRepository) InsertVendor(vendor models.Vendor) (*models.Vendor, error) {
	err := repo.DB.Create(&vendor).Error
	if err != nil {
		return nil, err
	}

	return &vendor, nil
}

func (repo *VendorRepository) CheckVendorExists(vendor_id int64) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.Vendor{}).Where("id = ?", vendor_id).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *VendorRepository) UpdateAverageRating(vendor_id int64) error {
	var average_rating float64

	err := repo.DB.
		Model(&models.Rating{}).
		Select("AVG(score)").
		Where("vendor_id = ?", vendor_id).
		Scan(&average_rating).Error
	if err != nil {
		return err
	}

	err = repo.DB.
		Model(&models.Vendor{}).
		Where("id = ?", vendor_id).
		Update("average_rating", average_rating).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *VendorRepository) GetVendorById(vendor_id int64) (*models.Vendor, error) {
	var vendor models.Vendor

	err := repo.DB.First(&vendor, vendor_id).Error
	if err != nil {
		return nil, err
	}

	return &vendor, nil
}
