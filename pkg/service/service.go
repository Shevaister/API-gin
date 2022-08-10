package service

import (
	"API"
	"API/pkg/repository"
)

type Authorization interface {
	CreateUser(user API_gin.Users) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetAllComments(postId int) ([]API_gin.Comments, error)
	GetCommentById(postId, commentId int) (API_gin.Comments, error)
	UpdateComment(commentId int, input API_gin.UpdateCommentInput) error
	DeleteComment(commentId int) error
}

type Service struct {
	Authorization
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
		Comment:       NewCommentService(repos.Comment, repos.Post),
	}
}
