package repository

import (
	"API"
	"gorm.io/gorm"
)

type Authorisation interface {
	CreateUser(user API.Users) (int, error)
	GetUser(email, password string) API.Users
}

type Post interface {
}

type Comment interface {
}

type Repository struct {
	Authorisation
	Post
	Comment
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthSQL(db),
	}
}
