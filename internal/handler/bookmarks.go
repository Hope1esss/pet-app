package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// addPetInBookmarksById godoc
//
//	@Summary    Add Pet to Bookmarks
//	@Security    ApiKeyAuth
//	@Tags      Bookmarks
//	@Description  Add a pet to the user's bookmarks by pet ID
//	@ID        add-pet-in-bookmarks
//	@Accept      json
//	@Produce    json
//	@Param      id    path    int  true  "Pet ID"
//	@Success    200    {object}  model.Pet
//	@Failure    400    {object}  ErrorResponse
//	@Failure    401    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /api/bookmarks/{id} [post]
func (h *Handler) addPetInBookmarksById(c *gin.Context) {
	petId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Id Not Found")
		return
	}

	userId := int64(id)
	pet, err := h.services.Bookmarks.AddPetInBookmarksById(userId, petId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pet)
}

// deletePetFromBookmarksById godoc
//
//	@Summary    Delete Pet from Bookmarks
//	@Security    ApiKeyAuth
//	@Tags      Bookmarks
//	@Description  Delete a pet from the user's bookmarks by pet ID
//	@ID        delete-pet-from-bookmarks
//	@Accept      json
//	@Produce    json
//	@Param      id    path    int  true  "Pet ID"
//	@Success    200    {object}  map[string]interface{}
//	@Failure    400    {object}  ErrorResponse
//	@Failure    401    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /api/bookmarks/{id} [delete]
func (h *Handler) deletePetFromBookmarksById(c *gin.Context) {
	petId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Id Not Found")
		return
	}
	userId := int64(id)
	err = h.services.Bookmarks.DeletePetFromBookmarksById(userId, petId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Pet deleted successfully",
	})
}

// getAllBookmarks godoc
//
//	@Summary    Get All Bookmarks
//	@Security    ApiKeyAuth
//	@Tags      Bookmarks
//	@Description  Get all bookmarks for the user
//	@ID        get-all-bookmarks
//	@Accept      json
//	@Produce    json
//	@Success    200    {array}    model.Pet
//	@Failure    400    {object}  ErrorResponse
//	@Failure    401    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /api/bookmarks [get]
func (h *Handler) getAllBookmarks(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Id Not Found")
		return
	}
	userId := int64(id)
	msg, err := h.services.Bookmarks.GetAllBookmarks(userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, msg)
}
