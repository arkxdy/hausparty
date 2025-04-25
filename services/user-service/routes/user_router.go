package routes

import (
	"hauparty/services/user-service/di"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, c *di.Container) {
	router.GET("/", app)
	router.GET("/health", healthCheck)
	v1 := router.Group("/user")
	{
		v1.GET("/")
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func app(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"app": "user-service"})
}
