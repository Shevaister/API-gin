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

	err := r.db.Create(&field).Select(field, "user", "title", "body").Last(&field).Error

	return field.Id, err
}

func (r *PostSQL) GetAll(userId int) ([]API.Posts, error) {
	var posts []API.Posts

	err := r.db.Find(&posts, userId).Error

	return posts, err
}

func (r *PostSQL) GetById(userId, postId int) (API.Posts, error) {
	var post API.Posts

	err := r.db.Find(&post, userId, postId).Error

	return post, err
}

func (r *PostSQL) Delete(userId, postId int) error {
	var post API.Posts

	err := r.db.Find(&post, userId, postId).Delete(&post).Error

	return err
}
