package service

import (
	"API"
	"API/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(userId int, post API.Posts) (int, error) {
	return s.repo.Create(userId, post)
}

func (s *PostService) GetAll(userId int) ([]API.Posts, error) {
	return s.repo.GetAll(userId)
}

func (s *PostService) GetById(userId, postId int) (API.Posts, error) {
	return s.repo.GetById(userId, postId)
}

func (s *PostService) Delete(userId, postId int) error {
	return s.repo.Delete(userId, postId)
}
