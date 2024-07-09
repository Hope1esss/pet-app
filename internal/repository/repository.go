package repository

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Pet interface {
	AddPet(pet model.Pet) (int, error)
	GetAllPets() ([]model.Pet, error)
	GetPetById(id int) (model.Pet, error)
	FindByBreed(breed string) ([]model.Pet, error)
	FindByType(petType string) ([]model.Pet, error)
}

type Bookmarks interface {
}
type Repository struct {
	Authorization
	Pet
	Bookmarks
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewUserPostgres(db),
		Pet:           NewPetPostgres(db),
		Bookmarks:     NewBookmarksPostgres(db),
	}
}
