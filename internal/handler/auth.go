package handler

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signUp godoc
//
//	@Summary    Sign Up
//	@Tags      Account
//	@Description  Сreates a new user in the system
//	@ID        sign-up
//	@Accept      json
//	@Produce    json
//	@Param      query  body  SignUpInput  true  "Account info"
//	@Success    200    {object}  map[string]interface{}
//	@Failure    400    {object}  ErrorResponse
//	@Router      /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signIn godoc
//
//	@Summary    Sign In
//	@Tags      Account
//	@Description  Sign In
//	@ID        sign-in
//	@Accept      json
//	@Produce    json
//	@Param      query  body    SignInInput  true  "Account info"
//	@Success    200    {object}  map[string]interface{}
//	@Failure    400    {object}  ErrorResponse
//	@Failure    500    {object}  ErrorResponse
//	@Router      /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
