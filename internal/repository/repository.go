package repository

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Pet interface {
}

type Repository struct {
	Authorization
	Pet
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewUserPostgres(db),
	}
}
