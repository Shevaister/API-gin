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
	Create(userId int, post API.Posts) (int, error)
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
		Post:          NewPostSQL(db),
	}
}
