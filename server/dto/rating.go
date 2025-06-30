package dto

import (
	"main/models"
)

type AddRatingRequest struct {
	Score  int64  `json:"score" binding:"required,gte=0,lte=10"`
	Review string `json:"review" binding:"required,max=191"`
}

type Rating struct {
	ID       int64  `json:"id"`
	VendorId int64  `json:"vendor_id"`
	Score    int64  `json:"score"`
	Review   string `json:"review"`
}

func ConvertRatingModelToDto(model *models.Rating) *Rating {
	return &Rating{
		ID:       model.ID,
		VendorId: model.VendorId,
		Score:    model.Score,
		Review:   model.Review,
	}
}
