package repository

import (
	"API"
	"gorm.io/gorm"
)

type CommentSQL struct {
	db *gorm.DB
}

func NewCommentSQL(db *gorm.DB) *CommentSQL {
	return &CommentSQL{db: db}
}

func (r *CommentSQL) CreateComment(userId, postId int, comment API.Comments) (int, error) {
	var user API.Users

	r.db.Table("users").Take(&user).Where(userId)

	field := API.Comments{Post: postId, Name: comment.Name,
		Email: user.Email, Body: comment.Body}

	err := r.db.Create(&field).Error
	r.db.Select(field, "postId", "name", "email", "body").Last(&field)

	return field.Id, err
}

func (r *CommentSQL) GetAllComments(postId int) ([]API.Comments, error) {
	var comments []API.Comments

	err := r.db.Where(API.Comments{Post: postId}).Find(&comments).Error

	return comments, err
}

func (r *CommentSQL) GetCommentById(postId, commentId int) (API.Comments, error) {
	var comment API.Comments

	err := r.db.First(&comment, postId, commentId).Error

	return comment, err
}

func (r *CommentSQL) DeleteComment(commentId int) error {
	var comment API.Comments

	err := r.db.Where(API.Comments{Id: commentId}).Delete(&comment).Error

	return err
}

func (r *CommentSQL) UpdateComment(commentId int, input API.UpdateCommentInput) error {
	err := r.db.Where(API.Comments{Id: commentId}).Updates(API.Comments{Body: *input.Body}).Error

	return err
}
