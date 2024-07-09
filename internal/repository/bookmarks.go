package repository

import "gorm.io/gorm"

type BookmarksPostgres struct {
	db *gorm.DB
}

func NewBookmarksPostgres(db *gorm.DB) *BookmarksPostgres {
	return &BookmarksPostgres{db: db}
}
