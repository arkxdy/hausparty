package repository

import (
	"context"
	models "hauparty/libs/db/models/ratings"

	"gorm.io/gorm"
)

type RatingRepository interface {
	CreateRating(ctx context.Context, rating *models.Rating) error
	GetPartyRatings(ctx context.Context, partyID string) ([]models.Rating, error)
	UpdateHostResponse(ctx context.Context, ratingID string, response string) error
	DeleteRating(ctx context.Context, ratingID string) error
}

// GORM Implementation
type RatingRepositoryGorm struct {
	db *gorm.DB
}

// DeleteRating implements RatingRepository.
func (r *RatingRepositoryGorm) DeleteRating(ctx context.Context, ratingID string) error {
	panic("unimplemented")
}

// GetPartyRatings implements RatingRepository.
func (r *RatingRepositoryGorm) GetPartyRatings(ctx context.Context, partyID string) ([]models.Rating, error) {
	panic("unimplemented")
}

// UpdateHostResponse implements RatingRepository.
func (r *RatingRepositoryGorm) UpdateHostResponse(ctx context.Context, ratingID string, response string) error {
	panic("unimplemented")
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &RatingRepositoryGorm{db: db}
}

func (r *RatingRepositoryGorm) CreateRating(ctx context.Context, rating *models.Rating) error {
	return r.db.WithContext(ctx).Create(rating).Error
}
