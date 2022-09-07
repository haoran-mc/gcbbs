package route

import "github.com/gin-gonic/gin"

// RegisterBackedRoute 注册前台路由
func RegisterBackendRoute(engine *gin.Engine) {
	group := engine.Group("backend")
	group.Use(isAdmin)
}
