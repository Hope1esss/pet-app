package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %s", err.Error())
	}
	hashedPassword := sha1.New()
	hashedPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashedPassword.Sum([]byte(os.Getenv("SALT"))))
}
