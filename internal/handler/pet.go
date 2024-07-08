package handler

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type petInput struct {
	Name        string    `json:"name" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Breed       string    `json:"breed"`
	Age         string    `json:"age" binding:"required"`
	Size        string    `json:"size"`
	Gender      string    `json:"gender" binding:"required"`
	Description string    `json:"description"`
	AddDate     time.Time `json:"addDate"`
	AddedBy     string    `json:"addedBy"`
}

func (h *Handler) addPet(c *gin.Context) {
	var input petInput
	userId, _ := c.Get("userId")
	pet := model.Pet{
		Name:          input.Name,
		Type:          input.Type,
		Breed:         input.Breed,
		Age:           input.Age,
		Size:          input.Size,
		Gender:        input.Gender,
		Description:   input.Description,
		AddDate:       time.Now(),
		AddedByUserId: userId.(int),
	}

	if err := c.BindJSON(&pet); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Pet.AddPet(pet)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPet(c *gin.Context) {

}
func (h *Handler) updatePetInfo(c *gin.Context) {

}

func (h *Handler) uploadPetImage(c *gin.Context) {

}

func (h *Handler) findByBreed(c *gin.Context) {

}

func (h *Handler) findByType(c *gin.Context) {

}

func (h *Handler) deletePet(c *gin.Context) {

}
