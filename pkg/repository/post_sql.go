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

func (r *PostSQL) CreatePost(userId int, post API_gin.Posts) (int, error) {
	field := API_gin.Posts{User: userId, Title: post.Title, Body: post.Body}

	err := r.db.Create(&field).Error
	r.db.Select(field, "user", "title", "body").Last(&field)

	return field.Id, err
}

func (r *PostSQL) GetAllPosts(userId int) ([]API_gin.Posts, error) {
	var posts []API_gin.Posts

	err := r.db.Where(API_gin.Posts{User: userId}).Find(&posts).Error

	return posts, err
}

func (r *PostSQL) GetPostById(userId, postId int) (API_gin.Posts, error) {
	var post API_gin.Posts

	err := r.db.First(&post, userId, postId).Error

	return post, err
}

func (r *PostSQL) UpdatePost(userId, postId int, input API_gin.UpdatePostInput) error {
	err := r.db.Where(API_gin.Posts{User: userId, Id: postId}).Updates(API_gin.Posts{Title: *input.Title,
		Body: *input.Body}).Error

	return err
}

func (r *PostSQL) DeletePost(userId, postId int) error {
	var post API_gin.Posts

	err := r.db.Where(API_gin.Posts{User: userId, Id: postId}).Delete(&post).Error

	return err
}
