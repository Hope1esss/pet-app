package repository

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"gorm.io/gorm"
)

type PetPostgres struct {
	db *gorm.DB
}

func NewPetPostgres(db *gorm.DB) *PetPostgres {
	return &PetPostgres{db: db}
}

func (r *PetPostgres) AddPet(pet model.Pet) (int, error) {
	result := r.db.Table("pets").Create(&pet)
	if result.Error != nil {
		return 0, result.Error
	}
	return pet.Id, nil
}
