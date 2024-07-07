package repository

type Authorization interface {
}

type Pet interface {
}

type Repository struct {
	Authorization
	Pet
}

func NewRepository() *Repository {
	return &Repository{}
}
