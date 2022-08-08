package service

import (
	"API"
	"API/pkg/repository"
)

type CommentService struct {
	repo     repository.Comment
	postRepo repository.Post
}

func NewCommentService(repo repository.Comment, postRepo repository.Post) *CommentService {
	return &CommentService{repo: repo, postRepo: postRepo}
}

func (s *CommentService) CreateComment(userId, postId int, comment API.Comments) (int, error) {
	_, err := s.postRepo.GetPostById(userId, postId)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateComment(userId, postId, comment)
}

func (s *CommentService) GetAllComments(userId int) ([]API.Comments, error) {
	return s.repo.GetAllComments(userId)
}

func (s *CommentService) GetCommentById(postId, commentId int) (API.Comments, error) {
	return s.repo.GetCommentById(postId, commentId)
}

func (s *CommentService) UpdateComment(commentId int, input API.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateComment(commentId, input)
}

func (s *CommentService) DeleteComment(commentId int) error {
	return s.repo.DeleteComment(commentId)
}
