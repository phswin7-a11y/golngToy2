package infra

import (
	"fmt"
	"golangToy2/internal/domain"
	"log/slog"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 데이터베이스 초기화 및 연결
func InitDB(cfg *Config) (*gorm.DB, error) {
	// .env 파일 로드 (파일이 없어도 오류를 내지 않고 시스템 환경변수 사용)
	if err := godotenv.Load(); err != nil {
		slog.Info(".env 파일을 찾을 수 없습니다. 시스템 환경변수를 사용합니다.")
	}

	// 필수 설정 체크
	if cfg.DBUser == "" || cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("필수 DB 설정이 누락되었습니다")
	}

	// DSN (Data Source Name) 구성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 실패: %w", err)
	}

	slog.Info("데이터베이스 연결 성공", "host", cfg.DBHost, "port", cfg.DBPort, "name", cfg.DBName)

	// 자동 마이그레이션 (테이블 생성)
	err = db.AutoMigrate(&domain.User{}, &domain.Board{}, &domain.Comment{})
	if err != nil {
		return nil, fmt.Errorf("데이터베이스 마이그레이션 실패: %w", err)
	}
	
	slog.Info("데이터베이스 마이그레이션 완료")
	return db, nil
}
