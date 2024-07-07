package service

import "github.com/Hope1esss/pet-app/internal/repository"

type Authorization interface {
}

type Pet interface {
}

type Service struct {
	Authorization
	Pet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
