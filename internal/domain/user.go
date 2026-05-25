package domain

import (
	"time"
)

// User 사용자 테이블 모델
type User struct {
	UserIdx   uint      `gorm:"primaryKey;autoIncrement;column:user_idx" json:"userIdx"`
	UserID    string    `gorm:"type:varchar(50);unique;not null;column:user_id" json:"userId" binding:"required"`
	Username  string    `gorm:"type:varchar(50);not null;column:username" json:"username" binding:"required"`
	Email     string    `gorm:"type:varchar(100);not null;column:email" json:"email" binding:"required,email"`
	Password  string    `gorm:"type:varchar(255);not null;column:password" json:"password,omitempty" binding:"required"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	LastLogin *time.Time `gorm:"column:last_login" json:"lastLogin"`
}

// TableName 테이블 이름 설정
func (User) TableName() string {
	return "t_user"
}
