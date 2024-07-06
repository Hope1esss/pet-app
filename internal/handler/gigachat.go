package handler

import (
	"github.com/Hope1esss/pet-app/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GigaChat(c *gin.Context) {
	var inputMessage string
	if err := c.BindJSON(&inputMessage); err != nil {
		NewValidationResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	response := service.Giga(inputMessage)

	c.String(http.StatusOK, response)

}
