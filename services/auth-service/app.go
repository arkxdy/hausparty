package main

import (
	"context"
	"hauparty/services/auth-service/di"
	"hauparty/services/auth-service/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func InitApp() {
	ctx := context.Background()
	container, _ := di.BuildContainer(ctx)

	router := gin.Default()

	routes.RegisterRoutes(router, container)

	if err := router.Run(":" + os.Getenv("AUTH_SERVICE_PORT")); err != nil {
		panic("failed " + err.Error())
	}
}
