package repository

import (
	"errors"
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

func (r *PetPostgres) GetAllPets() ([]model.Pet, error) {
	var pets []model.Pet
	result := r.db.Table("pets").Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}
	return pets, nil
}

func (r *PetPostgres) GetPetById(id int) (model.Pet, error) {
	var pet model.Pet
	result := r.db.Table("pets").Where("id =?", id).First(&pet)
	if result.Error != nil {
		return model.Pet{}, result.Error
	}
	return pet, nil
}

func (r *PetPostgres) FindByBreed(breed string) ([]model.Pet, error) {
	var pets []model.Pet
	result := r.db.Table("pets").Where("breed =?", breed).Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(pets) == 0 {
		return nil, errors.New("breed not found")
	}
	return pets, nil
}

func (r *PetPostgres) FindByType(petType string) ([]model.Pet, error) {
	var pets []model.Pet
	result := r.db.Table("pets").Where("type =?", petType).Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(pets) == 0 {
		return nil, errors.New("type not found")
	}
	return pets, nil
}

func (r *PetPostgres) DeletePetById(id int) error {
	result := r.db.Table("pets").Where("id =?", id).Delete(&model.Pet{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r *PetPostgres) UpdatePetInfoById(petId int, pet model.Pet) (int, error) {
	result := r.db.Table("pets").Where("id =?", petId).Updates(pet)
	if result.Error != nil {
		return 0, result.Error
	}
	return petId, nil
}
