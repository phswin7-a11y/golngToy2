package comment

import (
	"golangToy2/internal/common"
	"golangToy2/internal/domain"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateComment godoc
// @Summary 댓글 등록
// @Description 새로운 댓글을 등록합니다.
// @Tags Comment
// @Accept json
// @Produce json
// @Param comment body domain.Comment true "댓글 정보"
// @Success 200 {object} common.Response
// @Router /comments [post]
func (h *Handler) CreateComment(ctx *gin.Context) {
	var c domain.Comment
	if err := ctx.ShouldBindJSON(&c); err != nil {
		slog.Warn("댓글 등록 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}

	if err := h.service.CreateComment(ctx.Request.Context(), &c); err != nil {
		slog.Error("댓글 등록 실패", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "댓글 등록 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "댓글이 성공적으로 등록되었습니다.", c))
}

// GetCommentsByBoard godoc
// @Summary 게시글별 댓글 목록 조회 (Pagination)
// @Description 특정 게시글의 댓글 목록을 페이징하여 조회합니다.
// @Tags Comment
// @Produce json
// @Param boardId path int true "게시글 ID"
// @Param page query int false "페이지 번호 (기본값: 1)"
// @Param size query int false "페이지 크기 (기본값: 10)"
// @Success 200 {object} common.Response
// @Router /boards/{boardId}/comments [get]
func (h *Handler) GetCommentsByBoard(ctx *gin.Context) {
	boardIdStr := ctx.Param("boardId")
	boardId, err := strconv.Atoi(boardIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 게시글 ID 형식입니다.", nil))
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pageRes, err := h.service.GetCommentsByBoard(ctx.Request.Context(), uint(boardId), page, size)
	if err != nil {
		slog.Error("댓글 목록 조회 실패", "boardId", boardId, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "댓글 목록 조회 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "댓글 목록 조회가 완료되었습니다.", pageRes))
}

// UpdateComment godoc
// @Summary 댓글 수정
// @Description 댓글 정보를 수정합니다.
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "댓글 ID"
// @Param comment body domain.Comment true "수정할 댓글 정보"
// @Success 200 {object} common.Response
// @Router /comments/{id} [put]
func (h *Handler) UpdateComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 ID 형식입니다.", nil))
		return
	}

	var c domain.Comment
	if err := ctx.ShouldBindJSON(&c); err != nil {
		slog.Warn("댓글 수정 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}
	c.CommentID = uint(id)

	if err := h.service.UpdateComment(ctx.Request.Context(), &c); err != nil {
		slog.Error("댓글 수정 실패", "id", id, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "댓글 수정 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "댓글 정보가 수정되었습니다.", c))
}

// DeleteComment godoc
// @Summary 댓글 삭제
// @Description 댓글을 삭제합니다.
// @Tags Comment
// @Produce json
// @Param id path int true "댓글 ID"
// @Success 200 {object} common.Response
// @Router /comments/{id} [delete]
func (h *Handler) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 ID 형식입니다.", nil))
		return
	}

	if err := h.service.DeleteComment(ctx.Request.Context(), uint(id)); err != nil {
		slog.Error("댓글 삭제 실패", "id", id, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "댓글 삭제 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "댓글이 삭제되었습니다.", nil))
}
