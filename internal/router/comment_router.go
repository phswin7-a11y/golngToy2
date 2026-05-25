package router

import (
	"golangToy2/internal/comment"
	"github.com/gin-gonic/gin"
)

// CommentRouter 댓글 관련 라우팅 설정
func CommentRouter(rg *gin.RouterGroup, handler *comment.Handler) {
	comments := rg.Group("/comments")
	{
		comments.POST("", handler.CreateComment)
		comments.PUT("/:id", handler.UpdateComment)
		comments.DELETE("/:id", handler.DeleteComment)
	}
}
