package model

type User struct {
	Id                int    `json:"id"`
	GoogleId          string `json:"googleId"`
	Username          string `json:"username"`
	BookmarkedPetsIds []int  `json:"bookmarks"`
}
