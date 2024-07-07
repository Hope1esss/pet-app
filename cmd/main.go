package main

import (
	"context"
	"github.com/Hope1esss/pet-app/internal/handler"
	"github.com/Hope1esss/pet-app/internal/repository"
	"github.com/Hope1esss/pet-app/internal/service"
	"github.com/spf13/viper"
	"log"
	"net/http"
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

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка при инициализации конфига: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
