package router

import (
	"golangToy2/internal/board"
	"golangToy2/internal/comment"
	"github.com/gin-gonic/gin"
)

// BoardRouter 게시판 관련 라우팅 설정
func BoardRouter(rg *gin.RouterGroup, boardHandler *board.Handler, commentHandler *comment.Handler) {
	boards := rg.Group("/boards")
	{
		boards.POST("", boardHandler.CreateBoard)
		boards.GET("", boardHandler.GetBoards)
		boards.GET("/:id", boardHandler.GetBoard)
		boards.PUT("/:id", boardHandler.UpdateBoard)
		boards.DELETE("/:id", boardHandler.DeleteBoard)
		
		// 게시글별 댓글 조회 라우트
		boards.GET("/:boardId/comments", commentHandler.GetCommentsByBoard)
	}
}
