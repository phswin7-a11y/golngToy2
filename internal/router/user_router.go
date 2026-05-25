package router

import (
	"golangToy2/internal/user"
	"github.com/gin-gonic/gin"
)

// UserRouter 사용자 관련 라우팅 설정
func UserRouter(rg *gin.RouterGroup, handler *user.Handler) {
	users := rg.Group("/users")
	{
		users.POST("", handler.RegisterUser)
		users.GET("", handler.GetUsers)
		users.GET("/:idx", handler.GetUser)
		users.PUT("/:idx", handler.UpdateUser)
		users.DELETE("/:idx", handler.DeleteUser)
	}
}
