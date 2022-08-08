package service

import (
	"API"
	"API/pkg/repository"
)

type Authorization interface {
	CreateUser(user API.Users) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetAllComments(postId int) ([]API.Comments, error)
	GetCommentById(postId, commentId int) (API.Comments, error)
	UpdateComment(commentId int, input API.UpdateCommentInput) error
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
