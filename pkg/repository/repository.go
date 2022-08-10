package repository

import (
	"API"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user API_gin.Users) (int, error)
	GetUser(email, password string) (API_gin.Users, error)
}

type Post interface {
	CreatePost(userId int, post API_gin.Posts) (int, error)
	GetAllPosts(userId int) ([]API_gin.Posts, error)
	GetPostById(userId, postId int) (API_gin.Posts, error)
	UpdatePost(userId, postId int, input API_gin.UpdatePostInput) error
	DeletePost(userId, postId int) error
}

type Comment interface {
	CreateComment(userId, postId int, comment API_gin.Comments) (int, error)
	GetAllComments(userId int) ([]API_gin.Comments, error)
	GetCommentById(postId, commentId int) (API_gin.Comments, error)
	UpdateComment(commentId int, input API_gin.UpdateCommentInput) error
	DeleteComment(commentId int) error
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
		Comment:       NewCommentSQL(db),
	}
}
