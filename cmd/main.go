package main

import (
	"context"
	"github.com/Hope1esss/pet-app/internal/config"
	"github.com/Hope1esss/pet-app/internal/handler"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(c context.Context) error {
	return s.httpServer.Shutdown(c)
}

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
	srv := new(Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), app.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err.Error())
	}
}
