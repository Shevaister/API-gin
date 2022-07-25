package service

import "API/pkg/repository"

type Authorisation interface {
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
	return &Service{}
}
