package dto

import (
	"main/models"
)

type Vendor struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	AverageRating float64 `json:"average_rating"`
}

type AddVendorRequest struct {
	Name        string `json:"name" binding:"required,max=191"`
	Description string `json:"description" binding:"required,max=191"`
}

func ConvertVendorModelToDto(model *models.Vendor) *Vendor {
	return &Vendor{
		ID:            model.ID,
		Name:          model.Name,
		Description:   model.Description,
		AverageRating: model.AverageRating,
	}
}
