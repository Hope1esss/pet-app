package service

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
)

type PetService struct {
	repos repository.Pet
}

func NewPetService(repos repository.Pet) *PetService {
	return &PetService{repos: repos}
}
func (s *PetService) AddPet(pet model.Pet) (int, error) {
	return s.repos.AddPet(pet)
}
