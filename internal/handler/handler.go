package handler

import (
	"github.com/Hope1esss/pet-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("google-auth", h.googleAuth)
	}
	api := router.Group("/api")
	{
		pets := api.Group("/pets")
		{
			pets.POST("/", h.addPet)                        // Добавление животного
			pets.GET("/:id", h.getPet)                      // Получение данных животного
			pets.PUT("/:id", h.updatePetInfo)               // Изменение данных животного
			pets.POST("/:id/uploadImage", h.uploadPetImage) // Загрузка изображения животного
			pets.GET("/findByBreed", h.findByBreed)         // Поиск животных по породе
			pets.GET("/findByType", h.findByType)           // Поиск животных по типу животного (кошка, собака и т.д)
			pets.DELETE("/:id", h.deletePet)                // Удаление животного
		}

	}

	return router
}
