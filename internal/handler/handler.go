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
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	user := router.Group("/user", h.userIdentity)
	{
		user.POST("/addPetInBookmarks/:id", h.addPetInBookmarksById)             // Добавление животного в закладки по id
		user.DELETE("/deletePetFromBookmarks/:id", h.deletePetFromBookmarksById) // Удаление животного из закладок по id
		user.DELETE("/deleteUser", h.deleteUser)
	}

	api := router.Group("/api", h.userIdentity)
	{
		pets := api.Group("/pets")
		{
			pets.POST("/", h.addPet)                        // Добавление животного
			pets.GET("/:id", h.getPet)                      // Получение данных животного по id
			pets.PATCH("/:id", h.updatePetInfo)             // Изменение данных животного по id
			pets.POST("/:id/uploadImage", h.uploadPetImage) // Загрузка изображения животного по id
			pets.GET("/findByBreed", h.findByBreed)         // Поиск животных по породе
			pets.GET("/findByType", h.findByType)           // Поиск животных по типу животного (кошка, собака и т.д)
			pets.DELETE("/:id", h.deletePet)                // Удаление животного по id
		}
	}

	return router
}
