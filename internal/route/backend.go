package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/app/backend"
)

// RegisterBackedRoute 注册前台路由
func RegisterBackendRoute(engine *gin.Engine) {
	group := engine.Group("backend")
	group.Use(isAdmin)

	group.GET("/", backend.Home.IndexPage)
	group.GET("/users", backend.User.IndexPage)

	group.GET("/topics", backend.Topic.IndexPage)
	group.POST("/topics/:id", backend.Topic.DeleteSubmit)

	group.GET("notices", backend.Notice.IndexPage)
	group.GET("notices/publish", backend.Notice.PublishPage)
	group.POST("notices/publish", backend.Notice.PublishSubmit)
}
