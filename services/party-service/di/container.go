package di

import (
	"context"
	"fmt"
	"hauparty/services/party-service/handlers"
	"hauparty/services/party-service/services"
	"hausparty/libs/db"
	"hausparty/libs/db/repository"
	"os"
)

type Container struct {
	PartyHandler *handlers.PartyHandler
}

func BuildContainer(ctx context.Context) (*Container, error) {
	serviceName := os.Getenv("SERVICE_NAME")
	println("Starting service: %s", serviceName)

	dbFactory, err := db.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}

	if serviceName == "party" {
		if err := dbFactory.AutoMigrateParty(); err != nil {
			return nil, fmt.Errorf("migration: %w", err)
		}
	}

	db := dbFactory.GetPostgres()
	repo := repository.NewPartyRepository(db) // returns UserRepository interface
	service := services.NewPartyService(repo) // takes interface, not concrete
	handler := handlers.NewPartyHandler(service)

	return &Container{PartyHandler: handler}, nil
}
