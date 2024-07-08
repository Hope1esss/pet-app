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
}

type Repository struct {
	Authorization
	Pet
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewUserPostgres(db),
		Pet:           NewPetPostgres(db),
	}
}
