package config

import (
	"github.com/Hope1esss/pet-app/internal/repository"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitConfig() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
		ServerPort: os.Getenv("SERVER_PORT"),
	})

	if err != nil {
		log.Fatalf("Error initialazing database:  %s", err.Error())
	}

	return db
}
