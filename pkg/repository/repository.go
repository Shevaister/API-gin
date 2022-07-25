package repository

import "gorm.io/gorm"

type Authorisation interface {
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
	return &Repository{}
}
