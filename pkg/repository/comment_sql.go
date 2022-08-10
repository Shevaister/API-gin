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

func (r *CommentSQL) CreateComment(userId, postId int, comment API_gin.Comments) (int, error) {
	var user API_gin.Users

	r.db.Table("users").Take(&user).Where(userId)

	field := API_gin.Comments{Post: postId, Name: comment.Name,
		Email: user.Email, Body: comment.Body}

	err := r.db.Create(&field).Error
	r.db.Select(field, "postId", "name", "email", "body").Last(&field)

	return field.Id, err
}

func (r *CommentSQL) GetAllComments(postId int) ([]API_gin.Comments, error) {
	var comments []API_gin.Comments

	err := r.db.Where(API_gin.Comments{Post: postId}).Find(&comments).Error

	return comments, err
}

func (r *CommentSQL) GetCommentById(postId, commentId int) (API_gin.Comments, error) {
	var comment API_gin.Comments

	err := r.db.First(&comment, postId, commentId).Error

	return comment, err
}

func (r *CommentSQL) DeleteComment(commentId int) error {
	var comment API_gin.Comments

	err := r.db.Where(API_gin.Comments{Id: commentId}).Delete(&comment).Error

	return err
}

func (r *CommentSQL) UpdateComment(commentId int, input API_gin.UpdateCommentInput) error {
	err := r.db.Where(API_gin.Comments{Id: commentId}).Updates(API_gin.Comments{Body: *input.Body}).Error

	return err
}
