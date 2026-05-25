package common

import (
	"time"
)

// Response 공통 응답 구조체
type Response struct {
	Success   bool        `json:"success" example:"true"`
	Message   string      `json:"message" example:"요청이 성공적으로 처리되었습니다."`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp" example:"2026-05-11T17:20:51.8852115"`
}

// NewResponse 성공 응답 생성
func NewResponse(success bool, message string, data interface{}) Response {
	return Response{
		Success:   success,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format("2006-01-02T15:04:05.9999999"),
	}
}

// PageResponse 페이징 응답 구조체
type PageResponse struct {
	List       interface{} `json:"list"`
	TotalCount int64       `json:"totalCount" example:"100"`
	Page       int         `json:"page" example:"1"`
	Size       int         `json:"size" example:"10"`
	TotalPage  int         `json:"totalPage" example:"10"`
}
