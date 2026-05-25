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
	return r.db.Create(board).Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*domain.Board, error) {
	var board domain.Board
	err := r.db.Preload("User").First(&board, id).Error
	return &board, err
}

func (r *repository) Update(ctx context.Context, board *domain.Board) error {
	// Select를 사용하여 Title과 Content 필드를 명시적으로 업데이트 대상에 포함시킵니다.
	// 이렇게 하면 빈 문자열("")도 데이터베이스에 정상적으로 업데이트됩니다.
	if err := r.db.WithContext(ctx).Model(board).
		Select("Title", "Content").
		Updates(board).Error; err != nil {
		return fmt.Errorf("게시글 수정 실패: %w", err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(&domain.Board{}, id).Error
}

func (r *repository) FindAll(ctx context.Context, page, size int) ([]domain.Board, int64, error) {
	var boards []domain.Board
	var total int64

	offset := (page - 1) * size

	err := r.db.Model(&domain.Board{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Offset(offset).Limit(size).Find(&boards).Error
	return boards, total, err
}
