package di

import (
	"context"
	"fmt"
	"hauparty/services/user-service/handlers"
	"hauparty/services/user-service/services"
	"hausparty/libs/db"
	"hausparty/libs/db/repository"
	"os"

	"gorm.io/gorm"
)

type Container struct {
	UserHandler *handlers.UserHandler
}

var DB *gorm.DB

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
	service := services.NewUserService(repo) // takes interface, not concrete
	handler := handlers.NewUserHandler(service)

	return &Container{UserHandler: handler}, nil
}
