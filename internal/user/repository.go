package user

import (
	"context"
	"fmt"
	"golangToy2/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, idx uint) (*domain.User, error)
	FindByUserID(ctx context.Context, userID string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, idx uint) error
	FindAll(ctx context.Context, page, size int) ([]domain.User, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("사용자 생성 실패: %w", err)
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, idx uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, idx).Error; err != nil {
		return nil, fmt.Errorf("사용자 조회 실패 (IDX: %d): %w", idx, err)
	}
	return &user, nil
}

func (r *repository) FindByUserID(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("사용자 ID로 조회 실패: %w", err)
	}
	return &user, nil
}

func (r *repository) Update(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("사용자 정보 업데이트 실패: %w", err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, idx uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.User{}, idx).Error; err != nil {
		return fmt.Errorf("사용자 삭제 실패: %w", err)
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context, page, size int) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64
	
	offset := (page - 1) * size
	
	err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("사용자 총 개수 조회 실패: %w", err)
	}
	
	err = r.db.WithContext(ctx).Offset(offset).Limit(size).Find(&users).Error
	if err != nil {
		return nil, 0, fmt.Errorf("사용자 목록 조회 실패: %w", err)
	}
	
	return users, total, nil
}
