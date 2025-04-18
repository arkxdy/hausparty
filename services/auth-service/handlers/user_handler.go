package handlers

import (
	"hauparty/services/auth-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Implementation here
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Implementation here
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
