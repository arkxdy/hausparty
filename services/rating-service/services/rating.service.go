package services

import (
	"context"
	models "hausparty/libs/db/models/ratings"
	"hausparty/libs/db/repository"
)

type RatingService struct {
	repo repository.RatingRepository
}

func NewRatingService(repo repository.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) AddRating(ctx context.Context, r *models.Rating) error {
	// add validation logic here
	return s.repo.CreateRating(ctx, r)
}
