package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
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
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %s", err.Error())
	}
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}
func generatePasswordHash(password string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %s", err.Error())
	}
	hashedPassword := sha1.New()
	hashedPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashedPassword.Sum([]byte(os.Getenv("SALT"))))
}
