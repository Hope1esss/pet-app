package main

import (
	"github.com/Hope1esss/pet-app"
	"github.com/Hope1esss/pet-app/internal/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(pet_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err.Error())
	}
}
