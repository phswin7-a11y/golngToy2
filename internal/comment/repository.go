package comment

import (
	"context"
	"fmt"
	"golangToy2/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	FindByID(ctx context.Context, id uint) (*domain.Comment, error)
	FindByBoardID(ctx context.Context, boardID uint, page, size int) ([]domain.Comment, int64, error)
	Update(ctx context.Context, comment *domain.Comment) error
	Delete(ctx context.Context, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, comment *domain.Comment) error {
	if err := r.db.WithContext(ctx).Create(comment).Error; err != nil {
		return fmt.Errorf("댓글 생성 실패: %w", err)
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (*domain.Comment, error) {
	var comment domain.Comment
	if err := r.db.WithContext(ctx).Preload("User").First(&comment, id).Error; err != nil {
		return nil, fmt.Errorf("댓글 조회 실패 (ID: %d): %w", id, err)
	}
	return &comment, nil
}

func (r *repository) FindByBoardID(ctx context.Context, boardID uint, page, size int) ([]domain.Comment, int64, error) {
	var comments []domain.Comment
	var total int64
	
	offset := (page - 1) * size
	
	query := r.db.WithContext(ctx).Model(&domain.Comment{}).Where("board_id = ?", boardID)
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("게시글 댓글 총 개수 조회 실패: %w", err)
	}
	
	if err := query.Preload("User").Offset(offset).Limit(size).Find(&comments).Error; err != nil {
		return nil, 0, fmt.Errorf("게시글 댓글 목록 조회 실패: %w", err)
	}
	return comments, total, nil
}

func (r *repository) Update(ctx context.Context, comment *domain.Comment) error {
	// Select를 사용하여 Content 필드를 명시적으로 업데이트 대상에 포함시킵니다.
	// 이렇게 하면 빈 문자열("")도 정상적으로 업데이트됩니다.
	if err := r.db.WithContext(ctx).Model(comment).
		Select("Content").
		Updates(comment).Error; err != nil {
		return fmt.Errorf("댓글 수정 실패: %w", err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Comment{}, id).Error; err != nil {
		return fmt.Errorf("댓글 삭제 실패: %w", err)
	}
	return nil
}
