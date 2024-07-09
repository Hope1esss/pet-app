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
		user.GET("/recommendations", h.getRecommendationsFromGigaChat) // Получение рекомендаций от гигачата
		user.PATCH("/", h.updateUser)                                  // Изменение данных пользователя
		user.DELETE("/", h.deleteUser)                                 // Удаление пользователя
	}

	api := router.Group("/api", h.userIdentity)
	{
		pets := api.Group("/pets")
		{
			pets.POST("/", h.addPet)                       // Добавление животного DONE
			pets.GET("/:id", h.getPetById)                 // Получение данных животного по id DONE
			pets.GET("/", h.getAllPets)                    // Получение всех животных DONE
			pets.PATCH("/:id", h.updatePetInfoById)        // Изменение данных животного по id DONE
			pets.GET("/findByBreed/:breed", h.findByBreed) // Поиск животных по породе DONE
			pets.GET("/findByType/:type", h.findByType)    // Поиск животных по типу животного (кошка, собака и т.д) DONE
			pets.DELETE("/:id", h.deletePetById)           // Удаление животного по id DONE
		}

		bookmarks := api.Group("/bookmarks")
		{
			bookmarks.GET("/", h.getAllBookmarks)                  //Получение всех закладок DONE
			bookmarks.POST("/:id", h.addPetInBookmarksById)        // Добавление животного в закладки по id DONE
			bookmarks.DELETE("/:id", h.deletePetFromBookmarksById) // Удаление животного из закладок по id DONE
		}
	}

	return router
}
