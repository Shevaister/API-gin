package service

import (
	"API"
	"API/pkg/repository"
)

type Authorisation interface {
	CreateUser(user API.Users) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Post interface {
}

type Comment interface {
}

type Service struct {
	Authorisation
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
	}
}
