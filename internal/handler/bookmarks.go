package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
		"message": "pet deleted successfully",
	})

}
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
