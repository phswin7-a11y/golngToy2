package domain

import (
	"time"
)

// Comment 댓글 테이블 모델
type Comment struct {
	CommentID uint      `gorm:"primaryKey;autoIncrement;column:comment_id" json:"commentId"`
	BoardID   uint      `gorm:"column:board_id;not null" json:"boardId" binding:"required"`
	Board     *Board    `gorm:"foreignKey:BoardID;references:BoardID" json:"board,omitempty"`
	UserIdx   uint      `gorm:"column:user_idx;not null" json:"userIdx" binding:"required"`
	User      *User     `gorm:"foreignKey:UserIdx;references:UserIdx" json:"user,omitempty"`
	Content   string    `gorm:"type:text;not null;column:content" json:"content" binding:"required"`
	MadeAt    time.Time `gorm:"column:made_at;autoCreateTime" json:"madeAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	Step      int       `gorm:"column:step;default:0" json:"step"`
}

// TableName 테이블 이름 설정
func (Comment) TableName() string {
	return "t_comment"
}
