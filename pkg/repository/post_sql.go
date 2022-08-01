package repository

import (
	"API"
	"gorm.io/gorm"
)

type PostSQL struct {
	db *gorm.DB
}

func NewPostSQL(db *gorm.DB) *PostSQL {
	return &PostSQL{db: db}
}

func (r *PostSQL) Create(userId int, post API.Posts) (int, error) {
	field := API.Posts{User: userId, Title: post.Title, Body: post.Body}
	r.db.Create(&field)
	r.db.Select(field, "user", "title", "body")
	r.db.Last(&field)

	return field.Id, nil
}
