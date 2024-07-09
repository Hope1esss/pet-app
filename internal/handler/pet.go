package handler

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	EditedBy    string    `json:"addedBy"`
}

type getResponse struct {
	Data []model.Pet `json:"data"`
}

func (h *Handler) addPet(c *gin.Context) {
	var input petInput
	userId, _ := c.Get("userId")
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Type = strings.ToUpper(input.Type)
	input.Breed = strings.ToUpper(input.Breed)

	pet := model.Pet{
		Name:           input.Name,
		Type:           input.Type,
		Breed:          input.Breed,
		Age:            input.Age,
		Size:           input.Size,
		Gender:         input.Gender,
		Description:    input.Description,
		AddDate:        time.Now(),
		EditedByUserId: userId.(int),
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

func (h *Handler) getAllPets(c *gin.Context) {

	pets, err := h.services.Pet.GetAllPets()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getResponse{
		Data: pets,
	})
}

func (h *Handler) getPetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	pet, err := h.services.Pet.GetPetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (h *Handler) updatePetInfoById(c *gin.Context) {
	var input petInput
	userId, _ := c.Get("userId")
	petId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Type = strings.ToUpper(input.Type)
	input.Breed = strings.ToUpper(input.Breed)
	pet := model.Pet{
		Name:           input.Name,
		Type:           input.Type,
		Breed:          input.Breed,
		Age:            input.Age,
		Size:           input.Size,
		Gender:         input.Gender,
		Description:    input.Description,
		AddDate:        time.Now(),
		EditedByUserId: userId.(int),
	}
	id, err := h.services.Pet.UpdatePetInfoById(petId, pet)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) findByBreed(c *gin.Context) {
	breed := strings.ToUpper(c.Param("breed"))

	pets, err := h.services.Pet.FindByBreed(breed)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getResponse{
		Data: pets,
	})

}

func (h *Handler) findByType(c *gin.Context) {
	petType := strings.ToUpper(c.Param("type"))

	pets, err := h.services.Pet.FindByType(petType)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getResponse{
		Data: pets,
	})

}

func (h *Handler) deletePetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.Pet.DeletePetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pet deleted successfully",
	})
}
