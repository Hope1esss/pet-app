package service

import (
	"github.com/Hope1esss/pet-app/internal/model"
	"github.com/Hope1esss/pet-app/internal/repository"
)

type BookmarksService struct {
	repos repository.Bookmarks
}

func NewBookmarksService(repos repository.Bookmarks) *BookmarksService {
	return &BookmarksService{repos: repos}
}

func (s *BookmarksService) AddPetInBookmarksById(userId, petId int64) (model.Pet, error) {
	return s.repos.AddPetInBookmarksById(userId, petId)
}

func (s *BookmarksService) GetAllBookmarks(userId int64) ([]model.Pet, error) {
	return s.repos.GetAllBookmarks(userId)
}

func (s *BookmarksService) DeletePetFromBookmarksById(userId, petId int64) error {
	return s.repos.DeletePetFromBookmarksById(userId, petId)
}
