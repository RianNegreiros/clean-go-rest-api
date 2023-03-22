package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("error fetching comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment is representation of a comment
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store interface defines all of the methods needs
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// Service is the struct on which all of the logic will be built on top of
type Service struct {
	Store Store
}

// NewService creates a new comment service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Getting comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCommet Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCommet)
	if err != nil {
		fmt.Println("Error updating comment")
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
