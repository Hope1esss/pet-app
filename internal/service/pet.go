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

func (s *PetService) GetAllPets() ([]model.Pet, error) {
	return s.repos.GetAllPets()
}

func (s *PetService) GetPetById(id int) (model.Pet, error) {
	return s.repos.GetPetById(id)
}

func (s *PetService) FindByBreed(breed string) ([]model.Pet, error) {
	return s.repos.FindByBreed(breed)
}

func (s *PetService) FindByType(petType string) ([]model.Pet, error) {
	return s.repos.FindByType(petType)
}
