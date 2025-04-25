package handlers

import (
	"hauparty/services/rating-service/services"
	models "hausparty/libs/db/models/ratings"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handlers/rating_handler.go
func (h *RatingHandler) AddRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ratingService.AddRating(c.Request.Context(), &rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save rating"})
		return
	}
	c.JSON(http.StatusCreated, rating)
}

type RatingHandler struct {
	ratingService services.RatingService
}

func NewRatingHandler(s services.RatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: s,
	}
}
