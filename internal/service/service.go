package service

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Pet interface {
}

type Service struct {
	Authorization
	Pet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
