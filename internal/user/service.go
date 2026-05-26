package user

import (
	"context"
	"fmt"
	"golangToy2/internal/common"
	"golangToy2/internal/domain"
	"math"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, userID, password string) (string, string, error)
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
	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("service - 비밀번호 해싱 실패: %w", err)
	}
	user.Password = string(hashedPassword)

	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("service - 사용자 등록 실패: %w", err)
	}
	return nil
}

func (s *service) Login(ctx context.Context, userID, password string) (string, string, error) {
	user, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return "", "", fmt.Errorf("service - 사용자 조회 실패: %w", err)
	}

	// 비밀번호 확인
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", fmt.Errorf("service - 비밀번호가 일치하지 않습니다")
	}

	// 토큰 생성
	accessToken, refreshToken, err := common.GenerateToken(user.UserIdx, user.UserID)
	if err != nil {
		return "", "", fmt.Errorf("service - 토큰 생성 실패: %w", err)
	}

	return accessToken, refreshToken, nil
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
