package repository

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"gorm.io/gorm"
)

type BookmarksPostgres struct {
	db *gorm.DB
}

func NewBookmarksPostgres(db *gorm.DB) *BookmarksPostgres {
	return &BookmarksPostgres{db: db}
}

func (r *BookmarksPostgres) AddPetInBookmarksById(userId, petId int64) (model.Pet, error) {
	var user model.User
	var pet model.Pet
	takeUser := r.db.Table("users").Where("id =?", userId).Take(&user)
	if takeUser.Error != nil {
		return model.Pet{}, takeUser.Error
	}
	user.BookmarkedPetsIds = append(user.BookmarkedPetsIds, petId)
	result := r.db.Table("users").Where("id =?", userId).Updates(&user)
	if result.Error != nil {
		return model.Pet{}, result.Error
	}
	takePet := r.db.Table("pets").Where("id =?", petId).Take(&pet)
	if takePet.Error != nil {
		return model.Pet{}, takePet.Error
	}
	return pet, nil
}

func (r *BookmarksPostgres) GetAllBookmarks(userId int64) ([]model.Pet, error) {
	var user model.User
	var pet model.Pet
	var pets []model.Pet
	takeUser := r.db.Table("users").Where("id =?", userId).Take(&user)
	if takeUser.Error != nil {
		return nil, takeUser.Error
	}

	for _, petId := range user.BookmarkedPetsIds {
		takePet := r.db.Table("pets").Where("id =?", petId).First(&pet)
		pets = append(pets, pet)
		pet = model.Pet{}
		if takePet.Error != nil {
		}
	}
	return pets, nil
}

func (r *BookmarksPostgres) DeletePetFromBookmarksById(userId, petId int64) error {
	var user model.User
	takeUser := r.db.Table("users").Where("id =?", userId).Take(&user)
	if takeUser.Error != nil {
		return takeUser.Error
	}
	for index, id := range user.BookmarkedPetsIds {
		if id == petId {
			j := index
			for j < len(user.BookmarkedPetsIds)-1 {
				user.BookmarkedPetsIds[j] = user.BookmarkedPetsIds[j+1]
				j++
			}
			user.BookmarkedPetsIds = user.BookmarkedPetsIds[:len(user.BookmarkedPetsIds)-1]
		}
	}
	result := r.db.Table("users").Where("id =?", userId).Updates(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
