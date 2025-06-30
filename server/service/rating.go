package service

import (
	"main/dto"
	"main/models"
	"main/repository"
)

type RatingService struct {
	RatingRepo *repository.RatingRepository
}

func NewRatingService(ratingRepo *repository.RatingRepository) *RatingService {
	return &RatingService{
		RatingRepo: ratingRepo,
	}
}

func (service *RatingService) CreateNewRating(rating_request *dto.AddRatingRequest, vendor_id int64) (*dto.Rating, error) {
	rating, err := service.RatingRepo.InsertRating(models.Rating{
		VendorId: vendor_id,
		Score:    rating_request.Score,
		Review:   rating_request.Review,
	})
	if err != nil {
		return nil, err
	}

	return dto.ConvertRatingModelToDto(rating), nil
}

func (service *RatingService) GetRatingById(rating_id int64, vendor_id int64) (*dto.Rating, error) {
	rating, err := service.RatingRepo.GetRatingById(rating_id, vendor_id)
	if err != nil {
		return nil, err
	}

	return dto.ConvertRatingModelToDto(rating), nil
}

func (service *RatingService) GetRatingsByVendorId(vendor_id int64) ([]*dto.Rating, error) {
	ratings, err := service.RatingRepo.GetRatingsByVendorId(vendor_id)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.Rating, len(ratings))
	for index, rating_model := range ratings {
		result[index] = dto.ConvertRatingModelToDto(rating_model)
	}

	return result, nil
}
