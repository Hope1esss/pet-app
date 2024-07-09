package handler

import (
	_ "github.com/Hope1esss/pet-app/docs"
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

// addPet godoc
//
//	@Summary    Add Pet
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Add a new pet
//	@ID        add-pet
//	@Accept      json
//	@Produce    json
//	@Param    query  body    petInput  true  "Pet info"
//	@Success    200    {object}  map[string]interface{}
//	@Failure    400    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /api/pets [post]
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

// getAllPets godoc
//
//	@Summary    Get All Pets
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Get all pets
//	@ID        get-all-pets
//	@Accept      json
//	@Produce    json
//	@Success    200  {object}  getResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets [get]
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

// getPetById godoc
//
//	@Summary    Get Pet By ID
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Get pet by ID
//	@ID        get-pet-by-id
//	@Accept      json
//	@Produce    json
//	@Param      id    path    int  true  "Pet ID"
//	@Success    200  {object}  model.Pet
//	@Failure    400  {object}  ErrorResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets/{id} [get]
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

// updatePetInfoById godoc
//
//	@Summary    Update Pet Info By ID
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Update pet info by ID
//	@ID        update-pet-info-by-id
//	@Accept      json
//	@Produce    json
//	@Param      id    path    int      true  "Pet ID"
//	@Param     query body    petInput  true  "Updated pet info"
//	@Success    200  {object}  map[string]interface{}
//	@Failure    400  {object}  ErrorResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets/{id} [patch]
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

// findByBreed godoc
//
//	@Summary    Find Pets By Breed
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Find pets by breed
//	@ID        find-by-breed
//	@Accept      json
//	@Produce    json
//	@Param      breed  path    string  true  "Pet breed"
//	@Success    200  {object}  getResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets/findByBreed/{breed} [get]
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

// findByType godoc
//
//	@Summary    Find Pets By Type
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Find pets by type
//	@ID        find-by-type
//	@Accept      json
//	@Produce    json
//	@Param      type  path    string  true  "Pet type"
//	@Success    200  {object}  getResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets/findByType/{type} [get]
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

// deletePetById godoc
//
//	@Summary    Delete Pet By ID
//	@Security    ApiKeyAuth
//	@Tags      Pets
//	@Description  Delete pet by ID
//	@ID        delete-pet-by-id
//	@Accept      json
//	@Produce    json
//	@Param      id    path    int    true  "Pet ID"
//	@Success    200  {object}  map[string]interface{}
//	@Failure    400  {object}  ErrorResponse
//	@Failure    500  {object}  ErrorResponse
//	@Router      /api/pets/{id} [delete]
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
