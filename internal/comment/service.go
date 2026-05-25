package comment

import (
	"context"
	"fmt"
	"golangToy2/internal/common"
	"golangToy2/internal/domain"
	"math"
)

type Service interface {
	CreateComment(ctx context.Context, comment *domain.Comment) error
	GetComment(ctx context.Context, id uint) (*domain.Comment, error)
	GetCommentsByBoard(ctx context.Context, boardID uint, page, size int) (*common.PageResponse, error)
	UpdateComment(ctx context.Context, comment *domain.Comment) error
	DeleteComment(ctx context.Context, id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateComment(ctx context.Context, comment *domain.Comment) error {
	if err := s.repo.Create(ctx, comment); err != nil {
		return fmt.Errorf("service - 댓글 생성 실패: %w", err)
	}
	return nil
}

func (s *service) GetComment(ctx context.Context, id uint) (*domain.Comment, error) {
	comment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service - 댓글 조회 실패: %w", err)
	}
	return comment, nil
}

func (s *service) GetCommentsByBoard(ctx context.Context, boardID uint, page, size int) (*common.PageResponse, error) {
	comments, total, err := s.repo.FindByBoardID(ctx, boardID, page, size)
	if err != nil {
		return nil, fmt.Errorf("service - 게시글 댓글 목록 조회 실패: %w", err)
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return &common.PageResponse{
		List:       comments,
		TotalCount: total,
		Page:       page,
		Size:       size,
		TotalPage:  totalPage,
	}, nil
}

func (s *service) UpdateComment(ctx context.Context, comment *domain.Comment) error {
	if err := s.repo.Update(ctx, comment); err != nil {
		return fmt.Errorf("service - 댓글 수정 실패: %w", err)
	}
	return nil
}

func (s *service) DeleteComment(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("service - 댓글 삭제 실패: %w", err)
	}
	return nil
}
