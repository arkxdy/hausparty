package di

import (
	"context"
	"fmt"
	"hauparty/services/rating-service/handlers"
	"hauparty/services/rating-service/services"
	"hausparty/libs/db"
	"hausparty/libs/db/repository"
	"os"
)

type Container struct {
	RatingHandler *handlers.RatingHandler
}

func BuildContainer(ctx context.Context) (*Container, error) {
	serviceName := os.Getenv("SERVICE_NAME")
	println("Starting service: %s", serviceName)

	dbFactory, err := db.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}

	repo := repository.NewMongoRatingRepository(dbFactory.GetMongo())
	service := services.NewRatingService(repo)
	handler := handlers.NewRatingHandler(*service)

	return &Container{
		RatingHandler: handler,
	}, nil
}
