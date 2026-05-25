package board

import (
	"context"
	"fmt"
	"golangToy2/internal/common"
	"golangToy2/internal/domain"
	"math"
)

type Service interface {
	CreateBoard(ctx context.Context, board *domain.Board) error
	GetBoard(ctx context.Context, id uint) (*domain.Board, error)
	GetBoards(ctx context.Context, page, size int) (*common.PageResponse, error)
	UpdateBoard(ctx context.Context, board *domain.Board) error
	DeleteBoard(ctx context.Context, id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateBoard(ctx context.Context, board *domain.Board) error {
	if err := s.repo.Create(ctx, board); err != nil {
		return fmt.Errorf("service - 게시글 생성 실패: %w", err)
	}
	return nil
}

func (s *service) GetBoard(ctx context.Context, id uint) (*domain.Board, error) {
	board, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service - 게시글 조회 실패: %w", err)
	}
	return board, nil
}

func (s *service) GetBoards(ctx context.Context, page, size int) (*common.PageResponse, error) {
	boards, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, fmt.Errorf("service - 게시글 목록 조회 실패: %w", err)
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return &common.PageResponse{
		List:       boards,
		TotalCount: total,
		Page:       page,
		Size:       size,
		TotalPage:  totalPage,
	}, nil
}

func (s *service) UpdateBoard(ctx context.Context, board *domain.Board) error {
	if err := s.repo.Update(ctx, board); err != nil {
		return fmt.Errorf("service - 게시글 수정 실패: %w", err)
	}
	return nil
}

func (s *service) DeleteBoard(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("service - 게시글 삭제 실패: %w", err)
	}
	return nil
}
