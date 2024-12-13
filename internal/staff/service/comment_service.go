package service

import (
	"context"
	"errors"
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
)

type CommentService interface {
	GetAllComments() ([]model.Comment, error)
	CreateComment(ctx context.Context, req NewCommentRequest) (*model.Comment, error)
	DeleteOwnComment(ctx context.Context, commentID int) error
	DeleteUserComment(ctx context.Context, commentID int) error
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) GetAllComments() ([]model.Comment, error) {
	return s.repo.GetAllComments()
}

func (s *commentService) CreateComment(ctx context.Context, req NewCommentRequest) (*model.Comment, error) {
	comment := &model.Comment{
		ArticleID: req.ArticleID,
		Username:  req.Username,
		Email:     req.Email,
		Comment:   req.Comment,
		ParentID:  req.ParentID,
		Status:    "pending",
	}
	return s.repo.CreateComment(comment)
}

func (s *commentService) DeleteOwnComment(ctx context.Context, commentID int) error {
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if comment.Username != ctx.Value("username") {
		return errors.New("not authorized to delete this comment")
	}
	return s.repo.DeleteComment(commentID)
}

func (s *commentService) DeleteUserComment(ctx context.Context, commentID int) error {
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if comment.ArticleID != ctx.Value("articleID") {
		return errors.New("not authorized to delete this comment")
	}
	return s.repo.DeleteComment(commentID)
}

// NewCommentRequest represents a request to create a new comment
type NewCommentRequest struct {
	ArticleID int    `json:"article_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Comment   string `json:"comment"`
	ParentID  *int   `json:"parent_id,omitempty"`
}
