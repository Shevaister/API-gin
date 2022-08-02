package repository

import (
	"API"
	"fmt"
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
	err := r.db.Create(&field).Error
	r.db.Select(field, "user", "title", "body")
	r.db.Last(&field)

	return field.Id, err
}

func (r *PostSQL) GetAll(userId int) ([]API.Posts, error) {
	var posts []API.Posts

	err := r.db.Raw(fmt.Sprintf("SELECT * FROM `%s` WHERE `user` = ?", postsTable), userId).Scan(&posts).Error

	return posts, err
}

func (r *PostSQL) GetById(userId, postId int) (API.Posts, error) {
	var post API.Posts

	err := r.db.Raw(fmt.Sprintf("SELECT * FROM `%s` WHERE `user` = ? AND `id` = ?", postsTable),
		userId, postId).Scan(&post).Error

	return post, err
}
