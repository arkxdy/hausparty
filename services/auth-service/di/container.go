package di

import (
	"hauparty/services/auth-service/handlers"
	"hauparty/services/auth-service/repository"
	"hauparty/services/auth-service/services"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
}

func BuildContainer() *Container {
	repo := repository.NewUserRepository("Test")
	service := services.NewAuthService(repo)
	handler := handlers.NewAuthHandler(service)

	return &Container{
		AuthHandler: handler,
	}
}
