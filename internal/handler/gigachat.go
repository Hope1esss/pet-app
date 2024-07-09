package handler

import (
	"github.com/Hope1esss/pet-app/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getRecommendationsFromGigaChat godoc
//
//	@Summary    Get Recommendations
//	@Security    ApiKeyAuth
//	@Tags      Recommendations
//	@Description  Get recommendations from GigaChat
//	@ID        get-recommendations
//	@Accept      json
//	@Produce    plain
//	@Param      message  body    string  true  "Input message"
//	@Success    200    {string}  string  "Response message"
//	@Failure    400    {object}  ErrorResponse
//	@Failure    401    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /api/recommendations [post]
func (h *Handler) getRecommendationsFromGigaChat(c *gin.Context) {
	var inputMessage string
	if err := c.BindJSON(&inputMessage); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	response := service.GetRecommendationsFromGigaChat(inputMessage)

	c.String(http.StatusOK, response)
}
