package service

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Pet interface {
	AddPet(pet model.Pet) (int, error)
	GetAllPets() ([]model.Pet, error)
	GetPetById(id int) (model.Pet, error)
	FindByBreed(breed string) ([]model.Pet, error)
	FindByType(petType string) ([]model.Pet, error)
	DeletePetById(id int) error
	UpdatePetInfoById(petId int, pet model.Pet) (int, error)
}

type Bookmarks interface {
	AddPetInBookmarksById(userId, petId int64) (model.Pet, error)
	GetAllBookmarks(userId int64) ([]model.Pet, error)
	DeletePetFromBookmarksById(userId, petId int64) error
}
type Service struct {
	Authorization
	Pet
	Bookmarks
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Pet:           NewPetService(repos.Pet),
		Bookmarks:     NewBookmarksService(repos.Bookmarks),
	}
}
