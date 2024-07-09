package main

import (
	"github.com/Hope1esss/pet-app/internal/api"
	"github.com/Hope1esss/pet-app/internal/config"
	"github.com/Hope1esss/pet-app/internal/handler"
	"log"
	"os"
)

//	@title			Pet-app API
//	@version		1.0
//	@description	API Server PetApp Application

//@host localhost:8000
//@BasePath /

//@securityDefinitions.apikey ApiKeyAuth
//@in Header
//@name Authorization

func main() {
	db := config.InitConfig()
	app := handler.InitApp(db)
	srv := new(api.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), app.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err.Error())
	}
}
