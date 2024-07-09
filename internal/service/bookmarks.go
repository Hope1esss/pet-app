package service

import "github.com/Hope1esss/pet-app/internal/repository"

type BookmarksService struct {
	repos repository.Bookmarks
}

func NewBookmarksService(repos repository.Bookmarks) *BookmarksService {
	return &BookmarksService{repos: repos}
}
