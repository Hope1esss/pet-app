package model

import "github.com/lib/pq"

type User struct {
	Id                int           `json:"-" gorm:"primary_key"`
	Name              string        `json:"name" binding:"required"`
	Username          string        `json:"username" binding:"required"`
	Password          string        `json:"password" binding:"required"`
	BookmarkedPetsIds pq.Int64Array `json:"bookmarks" gorm:"type:integer[]"`
}
