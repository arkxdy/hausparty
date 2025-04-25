package repository

import (
	"context"
	models "hausparty/libs/db/models/ratings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingRepository interface {
	CreateRating(ctx context.Context, rating *models.Rating) error
	GetPartyRatings(ctx context.Context, partyID string) ([]models.Rating, error)
	UpdateHostResponse(ctx context.Context, ratingID string, response string) error
	DeleteRating(ctx context.Context, ratingID string) error
}

type mongoRatingRepo struct {
	coll *mongo.Collection
}

// CreateRating implements RatingRepository.
func (m *mongoRatingRepo) CreateRating(ctx context.Context, rating *models.Rating) error {
	rating.ID = primitive.NewObjectID()
	rating.CreatedAt = time.Now()
	_, err := m.coll.InsertOne(ctx, rating)
	return err
}

// DeleteRating implements RatingRepository.
func (m *mongoRatingRepo) DeleteRating(ctx context.Context, ratingID string) error {
	panic("unimplemented")
}

// GetPartyRatings implements RatingRepository.
func (m *mongoRatingRepo) GetPartyRatings(ctx context.Context, partyID string) ([]models.Rating, error) {
	panic("unimplemented")
}

// UpdateHostResponse implements RatingRepository.
func (m *mongoRatingRepo) UpdateHostResponse(ctx context.Context, ratingID string, response string) error {
	panic("unimplemented")
}

func NewMongoRatingRepository(db *mongo.Database) RatingRepository {
	return &mongoRatingRepo{
		coll: db.Collection("ratings"),
	}
}
