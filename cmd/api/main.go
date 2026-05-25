package main

import (
	"golangToy2/internal/board"
	"golangToy2/internal/comment"
	"golangToy2/internal/infra"
	"golangToy2/internal/router"
	"golangToy2/internal/user"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	_ "golangToy2/docs" // Swagger docs
)

// @title Golang Toy Project API
// @version 1.0
// @description 이 프로젝트는 Golang, Gin, GORM을 이용한 토이 프로젝트 API입니다.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 설정 로드
	cfg := infra.LoadConfig()

	// 로거 설정 (JSON 포맷 사용 가능)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// DB 초기화
	db, err := infra.InitDB(cfg)
	if err != nil {
		slog.Error("데이터베이스 초기화 실패", "error", err)
		os.Exit(1)
	}

	// 의존성 주입 (User)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// 의존성 주입 (Board)
	boardRepo := board.NewRepository(db)
	boardService := board.NewService(boardRepo)
	boardHandler := board.NewHandler(boardService)

	// 의존성 주입 (Comment)
	commentRepo := comment.NewRepository(db)
	commentService := comment.NewService(commentRepo)
	commentHandler := comment.NewHandler(commentService)

	// Gin 엔진 설정
	r := gin.Default()

	// 라우터 설정 (분리된 라우터 파일 사용)
	handlers := router.RouterHandlers{
		UserHandler:    userHandler,
		BoardHandler:   boardHandler,
		CommentHandler: commentHandler,
	}
	router.SetupRouter(r, handlers)

	// 서버 실행
	slog.Info("서버 시작 중", "port", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		slog.Error("서버 실행 실패", "error", err)
		os.Exit(1)
	}
}
