package main

import (
	"hauparty/services/auth-service/di"
	"hauparty/services/auth-service/routes"

	"github.com/gin-gonic/gin"
)

func InitApp() {
	container := di.BuildContainer()

	router := gin.Default()

	routes.RegisterRoutes(router, container)

	if err := router.Run(":8001"); err != nil {
		panic("failed " + err.Error())
	}
	// go func() {
	// 	if err := router.Run(); err != nil {
	// 		panic("failed "+err.Error())
	// 	}
	// }()

	// select {}
}
