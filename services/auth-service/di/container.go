package di

import (
	"context"
	"fmt"
	"hauparty/services/auth-service/handlers"
	"hauparty/services/auth-service/services"
	"hausparty/libs/db"
	"hausparty/libs/db/repository"
	"os"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
}

func BuildContainer(ctx context.Context) (*Container, error) {
	serviceName := os.Getenv("SERVICE_NAME")
	println("Starting service: %s", serviceName)

	dbFactory, err := db.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}

	if serviceName == "identity" {
		if err := dbFactory.AutoMigrateIdentity(); err != nil {
			return nil, fmt.Errorf("migration: %w", err)
		}
	}

	db := dbFactory.GetPostgres()
	repo := repository.NewUserRepository(db) // returns UserRepository interface
	service := services.NewAuthService(repo) // takes interface, not concrete
	handler := handlers.NewAuthHandler(service)

	return &Container{AuthHandler: handler}, nil
}
