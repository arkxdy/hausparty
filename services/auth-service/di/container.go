package di

import (
	"hauparty/libs/db/repository"
	"hauparty/services/auth-service/handlers"
	"hauparty/services/auth-service/services"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
}

func BuildContainer() *Container {
	repo := repository.NewUserRepository(nil)
	service := services.NewAuthService(repo)
	handler := handlers.NewAuthHandler(service)

	return &Container{
		AuthHandler: handler,
	}
}
