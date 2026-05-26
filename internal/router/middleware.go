package router

import (
	"golangToy2/internal/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 인증 미들웨어
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, common.NewResponse(false, "인증 헤더가 없습니다.", nil))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, common.NewResponse(false, "인증 형식이 올바르지 않습니다.", nil))
			c.Abort()
			return
		}

		claims, err := common.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewResponse(false, "유효하지 않은 토큰입니다.", nil))
			c.Abort()
			return
		}

		// 사용자 정보를 컨텍스트에 저장
		c.Set("userIdx", claims.UserIdx)
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
