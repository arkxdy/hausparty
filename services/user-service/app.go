package main

import (
	"context"
	"hauparty/services/user-service/di"
	"hauparty/services/user-service/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func InitApp() {
	ctx := context.Background()
	container, err := di.BuildContainer(ctx)

	if err != nil {
		println("Error", err.Error())
	}

	router := gin.Default()

	routes.RegisterRoutes(router, container)

	if err := router.Run(":" + os.Getenv("USER_SERVICE_PORT")); err != nil {
		panic("failed " + err.Error())
	}
}
