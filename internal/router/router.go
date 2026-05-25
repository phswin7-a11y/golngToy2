package router

import (
	"golangToy2/internal/board"
	"golangToy2/internal/comment"
	"golangToy2/internal/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterHandlers 라우터에 주입할 핸들러 모음
type RouterHandlers struct {
	UserHandler    *user.Handler
	BoardHandler   *board.Handler
	CommentHandler *comment.Handler
}

// SetupRouter 전체 라우팅 설정
func SetupRouter(r *gin.Engine, handlers RouterHandlers) {
	// Swagger 설정
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 그룹
	v1 := r.Group("/api/v1")
	{
		UserRouter(v1, handlers.UserHandler)
		BoardRouter(v1, handlers.BoardHandler, handlers.CommentHandler)
		CommentRouter(v1, handlers.CommentHandler)
	}
}
