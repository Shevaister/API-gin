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
	Create(userId int, post API.Posts) (int, error)
	GetAll(userId int) ([]API.Posts, error)
	GetById(userId, postId int) (API.Posts, error)
}

type Comment interface {
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
	}
}
