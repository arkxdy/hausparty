package routes

import (
	"hauparty/services/auth-service/di"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, c *di.Container) {
	router.GET("/", app)
	router.GET("/health", healthCheck)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", c.AuthHandler.Register)
		v1.POST("/login", c.AuthHandler.Login)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func app(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"app": "auth-service"})
}
