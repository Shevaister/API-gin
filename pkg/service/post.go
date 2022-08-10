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

func (s *PostService) CreatePost(userId int, post API_gin.Posts) (int, error) {
	return s.repo.CreatePost(userId, post)
}

func (s *PostService) GetAllPosts(userId int) ([]API_gin.Posts, error) {
	return s.repo.GetAllPosts(userId)
}

func (s *PostService) GetPostById(userId, postId int) (API_gin.Posts, error) {
	return s.repo.GetPostById(userId, postId)
}

func (s *PostService) UpdatePost(userId, postId int, input API_gin.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdatePost(userId, postId, input)
}

func (s *PostService) DeletePost(userId, postId int) error {
	return s.repo.DeletePost(userId, postId)
}
