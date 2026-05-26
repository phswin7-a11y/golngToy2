package board

import (
	"context"
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
	return s.repo.Create(ctx, board)
}

func (s *service) GetBoard(ctx context.Context, id uint) (*domain.Board, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetBoards(ctx context.Context, page, size int) (*common.PageResponse, error) {
	boards, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, err
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
	return s.repo.Update(ctx, board)
}

func (s *service) DeleteBoard(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
