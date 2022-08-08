package repository

import (
	"API"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user API.Users) (int, error)
	GetUser(email, password string) (API.Users, error)
}

type Post interface {
	CreatePost(userId int, post API.Posts) (int, error)
	GetAllPosts(userId int) ([]API.Posts, error)
	GetPostById(userId, postId int) (API.Posts, error)
	UpdatePost(userId, postId int, input API.UpdatePostInput) error
	DeletePost(userId, postId int) error
}

type Comment interface {
	CreateComment(userId, postId int, comment API.Comments) (int, error)
	GetAllComments(userId int) ([]API.Comments, error)
	GetCommentById(postId, commentId int) (API.Comments, error)
	UpdateComment(commentId int, input API.UpdateCommentInput) error
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
