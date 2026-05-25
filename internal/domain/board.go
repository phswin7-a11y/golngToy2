package domain

import (
	"time"
)

// Board 게시판 테이블 모델
type Board struct {
	BoardID   uint      `gorm:"primaryKey;autoIncrement;column:board_id" json:"boardId"`
	UserIdx   uint      `gorm:"column:user_idx;not null" json:"userIdx" binding:"required"`
	User      User      `gorm:"foreignKey:UserIdx;references:UserIdx" json:"user,omitempty"`
	Title     string    `gorm:"type:varchar(200);not null;column:title" json:"title" binding:"required"`
	Content   string    `gorm:"type:text;not null;column:content" json:"content" binding:"required"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

// TableName 테이블 이름 설정
func (Board) TableName() string {
	return "t_board"
}
