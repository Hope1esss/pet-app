package model

import "time"

type Pet struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Breed       string    `json:"breed"`
	Age         string    `json:"age" binding:"required"`
	Size        string    `json:"size"`
	Gender      string    `json:"gender" binding:"required"`
	Description string    `json:"description"`
	AddDate     time.Time `json:"addDate"`
}
