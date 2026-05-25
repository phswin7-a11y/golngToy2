package user

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

// RegisterUser godoc
// @Summary 사용자 등록
// @Description 새로운 사용자를 등록합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.User true "사용자 정보"
// @Success 200 {object} common.Response
// @Router /users [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		slog.Warn("사용자 등록 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}

	if err := h.service.RegisterUser(ctx.Request.Context(), &user); err != nil {
		slog.Error("사용자 등록 실패", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "사용자 등록 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "사용자가 성공적으로 등록되었습니다.", user))
}

// GetUser godoc
// @Summary 사용자 조회
// @Description IDX를 통해 사용자를 조회합니다.
// @Tags User
// @Produce json
// @Param idx path int true "사용자 IDX"
// @Success 200 {object} common.Response
// @Router /users/{idx} [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	idxStr := ctx.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 IDX 형식입니다.", nil))
		return
	}

	user, err := h.service.GetUser(ctx.Request.Context(), uint(idx))
	if err != nil {
		slog.Error("사용자 조회 실패", "idx", idx, "error", err)
		ctx.JSON(http.StatusNotFound, common.NewResponse(false, "사용자를 찾을 수 없습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "사용자 조회가 완료되었습니다.", user))
}

// GetUsers godoc
// @Summary 사용자 목록 조회 (Pagination)
// @Description 전체 사용자 목록을 페이징하여 조회합니다.
// @Tags User
// @Produce json
// @Param page query int false "페이지 번호 (기본값: 1)"
// @Param size query int false "페이지 크기 (기본값: 10)"
// @Success 200 {object} common.Response
// @Router /users [get]
func (h *Handler) GetUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pageRes, err := h.service.GetUsers(ctx.Request.Context(), page, size)
	if err != nil {
		slog.Error("사용자 목록 조회 실패", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "사용자 목록 조회 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "사용자 목록 조회가 완료되었습니다.", pageRes))
}

// UpdateUser godoc
// @Summary 사용자 정보 수정
// @Description 사용자 정보를 수정합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param idx path int true "사용자 IDX"
// @Param user body domain.User true "수정할 사용자 정보"
// @Success 200 {object} common.Response
// @Router /users/{idx} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	idxStr := ctx.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 IDX 형식입니다.", nil))
		return
	}

	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		slog.Warn("사용자 수정 바인딩 실패", "error", err)
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 요청 형식입니다.", nil))
		return
	}
	user.UserIdx = uint(idx)

	if err := h.service.UpdateUser(ctx.Request.Context(), &user); err != nil {
		slog.Error("사용자 수정 실패", "idx", idx, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "사용자 수정 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "사용자 정보가 수정되었습니다.", user))
}

// DeleteUser godoc
// @Summary 사용자 삭제
// @Description 사용자를 삭제합니다.
// @Tags User
// @Produce json
// @Param idx path int true "사용자 IDX"
// @Success 200 {object} common.Response
// @Router /users/{idx} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	idxStr := ctx.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewResponse(false, "잘못된 IDX 형식입니다.", nil))
		return
	}

	if err := h.service.DeleteUser(ctx.Request.Context(), uint(idx)); err != nil {
		slog.Error("사용자 삭제 실패", "idx", idx, "error", err)
		ctx.JSON(http.StatusInternalServerError, common.NewResponse(false, "사용자 삭제 중 오류가 발생했습니다.", nil))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(true, "사용자가 삭제되었습니다.", nil))
}
