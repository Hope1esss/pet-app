package handler

import (
	"github.com/Hope1esss/pet-app/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getRecommendationsFromGigaChat(c *gin.Context) {
	var inputMessage string
	if err := c.BindJSON(&inputMessage); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	response := service.GetRecommendationsFromGigaChat(inputMessage)

	c.String(http.StatusOK, response)
}
