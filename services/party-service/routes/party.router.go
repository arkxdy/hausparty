package routes

import (
	"hauparty/services/party-service/di"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, c *di.Container) {
	router.GET("/", app)
	router.GET("/health", healthCheck)
	v1 := router.Group("/auth")
	{
		v1.GET("/:provider")
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func app(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"app": "party-service"})
}
