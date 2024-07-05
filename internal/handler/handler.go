package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.auth)
	}

	api := router.Group("/api")
	{
		api.GET("/GigaChat", h.GigaChat)
	}
	return router
}