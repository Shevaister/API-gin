package repository

import (
	"API"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user API.Users) (int, error)
	GetUser(email, password string) API.Users
}

type Post interface {
}

type Comment interface {
}

type Repository struct {
	Authorization
	Post
	Comment
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
	}
}
