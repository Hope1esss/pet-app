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
	DeletePetById(id int) error
	UpdatePetInfoById(petId int, pet model.Pet) (int, error)
}

type Bookmarks interface {
	AddPetInBookmarksById(userId, petId int64) (model.Pet, error)
	GetAllBookmarks(userId int64) ([]model.Pet, error)
	DeletePetFromBookmarksById(userId, petId int64) error
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
