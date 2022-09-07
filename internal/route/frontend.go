package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/app/frontend"
)

// RegisterFrontedRoute 注册后台路由
func RegisterFrontedRoute(engine *gin.Engine) {
	group := engine.Group("/")

	// 社区首页
	group.GET("/", frontend.Home.HomePage)

	// 用户注册
	group.GET("/register", frontend.Auth.RegisterPage)
}
