package board

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

// CreateBoard godoc
// @Summary 게시글 등록
// @Description 새로운 게시글을 등록합니다.
// @Tags Board
// @Accept json
// @Produce json
// @Param board body domain.Board true "게시글 정보"
// @Success 200 {object} common.Response
// @Router /boards [post]
func (h *Handler) CreateBoard(ctx *gin.Context) {
	var board domain.Board
	if err := ctx.ShouldBindJSON(&board); err != nil {
		slog.Warn("게시글 등록 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}

	if err := h.service.CreateBoard(ctx.Request.Context(), &board); err != nil {
		slog.Error("게시글 등록 실패", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "게시글 등록 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "게시글이 성공적으로 등록되었습니다.", board))
}

// GetBoard godoc
// @Summary 게시글 조회
// @Description ID를 통해 게시글을 조회합니다.
// @Tags Board
// @Produce json
// @Param id path int true "게시글 ID"
// @Success 200 {object} common.Response
// @Router /boards/{id} [get]
func (h *Handler) GetBoard(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 ID 형식입니다.", nil))
		return
	}

	board, err := h.service.GetBoard(ctx.Request.Context(), uint(id))
	if err != nil {
		slog.Error("게시글 조회 실패", "id", id, "error", err)
		ctx.JSON(http.StatusNotFound, common.NewResponse(false, "게시글을 찾을 수 없습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "게시글 조회가 완료되었습니다.", board))
}

// GetBoards godoc
// @Summary 게시글 목록 조회 (Pagination)
// @Description 전체 게시글 목록을 페이징하여 조회합니다.
// @Tags Board
// @Produce json
// @Param page query int false "페이지 번호 (기본값: 1)"
// @Param size query int false "페이지 크기 (기본값: 10)"
// @Success 200 {object} common.Response
// @Router /boards [get]
func (h *Handler) GetBoards(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pageRes, err := h.service.GetBoards(ctx.Request.Context(), page, size)
	if err != nil {
		slog.Error("게시글 목록 조회 실패", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "게시글 목록 조회 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "게시글 목록 조회가 완료되었습니다.", pageRes))
}

// UpdateBoard godoc
// @Summary 게시글 수정
// @Description 게시글 정보를 수정합니다.
// @Tags Board
// @Accept json
// @Produce json
// @Param id path int true "게시글 ID"
// @Param board body domain.Board true "수정할 게시글 정보"
// @Success 200 {object} common.Response
// @Router /boards/{id} [put]
func (h *Handler) UpdateBoard(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 ID 형식입니다.", nil))
		return
	}

	var board domain.Board
	if err := ctx.ShouldBindJSON(&board); err != nil {
		slog.Warn("게시글 수정 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}
	board.BoardID = uint(id)

	if err := h.service.UpdateBoard(ctx.Request.Context(), &board); err != nil {
		slog.Error("게시글 수정 실패", "id", id, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "게시글 수정 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "게시글 정보가 수정되었습니다.", board))
}

// DeleteBoard godoc
// @Summary 게시글 삭제
// @Description 게시글을 삭제합니다.
// @Tags Board
// @Produce json
// @Param id path int true "게시글 ID"
// @Success 200 {object} common.Response
// @Router /boards/{id} [delete]
func (h *Handler) DeleteBoard(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 ID 형식입니다.", nil))
		return
	}

	if err := h.service.DeleteBoard(ctx.Request.Context(), uint(id)); err != nil {
		slog.Error("게시글 삭제 실패", "id", id, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "게시글 삭제 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "게시글이 삭제되었습니다.", nil))
}
