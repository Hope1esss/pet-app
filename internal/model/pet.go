package model

import "net/url"

type Pet struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age"`
	Size        string  `json:"size"`
	Description string  `json:"description"`
	PhotoURL    url.URL `json:"photo_url"`
}
