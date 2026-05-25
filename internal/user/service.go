package user

import (
	"context"
	"fmt"
	"golangToy2/internal/common"
	"golangToy2/internal/domain"
	"math"
)

type Service interface {
	RegisterUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, idx uint) (*domain.User, error)
	GetUsers(ctx context.Context, page, size int) (*common.PageResponse, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, idx uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, user *domain.User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("service - 사용자 등록 실패: %w", err)
	}
	return nil
}

func (s *service) GetUser(ctx context.Context, idx uint) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, idx)
	if err != nil {
		return nil, fmt.Errorf("service - 사용자 조회 실패: %w", err)
	}
	return user, nil
}

func (s *service) GetUsers(ctx context.Context, page, size int) (*common.PageResponse, error) {
	users, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, fmt.Errorf("service - 사용자 목록 조회 실패: %w", err)
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return &common.PageResponse{
		List:       users,
		TotalCount: total,
		Page:       page,
		Size:       size,
		TotalPage:  totalPage,
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, user *domain.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("service - 사용자 수정 실패: %w", err)
	}
	return nil
}

func (s *service) DeleteUser(ctx context.Context, idx uint) error {
	if err := s.repo.Delete(ctx, idx); err != nil {
		return fmt.Errorf("service - 사용자 삭제 실패: %w", err)
	}
	return nil
}
