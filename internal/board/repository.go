package board

import (
	"context"
	"fmt"
	"golangToy2/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, board *domain.Board) error
	FindByID(ctx context.Context, id uint) (*domain.Board, error)
	Update(ctx context.Context, board *domain.Board) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context, page, size int) ([]domain.Board, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, board *domain.Board) error {
	if err := r.db.WithContext(ctx).Create(board).Error; err != nil {
		return fmt.Errorf("게시글 생성 실패: %w", err)
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (*domain.Board, error) {
	var board domain.Board
	if err := r.db.WithContext(ctx).Preload("User").First(&board, id).Error; err != nil {
		return nil, fmt.Errorf("게시글 조회 실패 (ID: %d): %w", id, err)
	}
	return &board, nil
}

func (r *repository) Update(ctx context.Context, board *domain.Board) error {
	if err := r.db.WithContext(ctx).Save(board).Error; err != nil {
		return fmt.Errorf("게시글 수정 실패: %w", err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Board{}, id).Error; err != nil {
		return fmt.Errorf("게시글 삭제 실패: %w", err)
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context, page, size int) ([]domain.Board, int64, error) {
	var boards []domain.Board
	var total int64
	
	offset := (page - 1) * size
	
	err := r.db.WithContext(ctx).Model(&domain.Board{}).Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("게시글 총 개수 조회 실패: %w", err)
	}
	
	err = r.db.WithContext(ctx).Preload("User").Offset(offset).Limit(size).Find(&boards).Error
	if err != nil {
		return nil, 0, fmt.Errorf("게시글 목록 조회 실패: %w", err)
	}
	
	return boards, total, nil
}
