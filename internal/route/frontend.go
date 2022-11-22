package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/app/frontend"
)

// RegisterFrontedRoute 注册后台路由
func RegisterFrontedRoute(engine *gin.Engine) {
	group := engine.Group("/")

	group.GET("/", frontend.Home.HomePage) // 社区首页

	group.GET("/register", frontend.Auth.RegisterPage)    // 用户注册
	group.POST("/register", frontend.Auth.RegisterSubmit) // 提交注册

	group.GET("/login", frontend.Auth.LoginPage)     // 用户登录
	group.POST("/login", frontend.Auth.LoginSubmit)  // 登录提交
	group.GET("/logout", frontend.Auth.LogoutSubmit) // 登出用户

	group.GET("/publish", frontend.Topic.PublishPage)    // 话题发布
	group.POST("/publish", frontend.Topic.PublishSubmit) // 话题提交
	group.GET("/topics/:id", frontend.Topic.DetailPage)  // 话题详情

	group.POST("/comments", frontend.Comment.PublishSubmit)       // 评论话题
	group.POST("/comments/delete", frontend.Comment.DeleteSubmit) // 删除评论

	group.GET("/user", frontend.User.HomePage)         // 用户中心
	group.GET("/user/edit", frontend.User.EditPage)    // 用户编辑
	group.POST("/user/edit", frontend.User.EditSubmit) // 编辑提交

	group.POST("/md-upload", frontend.File.MDUploadSubmit)

	group.GET("/notice", frontend.Notice.HomePage) // 用户通知

	group.POST("/likes", frontend.Like.LikeSubmit)        // 用户点赞
	group.POST("/follows", frontend.Follow.FollowSubmit)  // 用户关注
	group.POST("/checkins", frontend.Checkin.StoreSubmit) // 用户签到

	group.POST("/reports", frontend.Report.ReportSubmit) // 举报资源
}
