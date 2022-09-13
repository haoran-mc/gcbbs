package webserver

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/haoran-mc/gcbbs/internal/route"
	"github.com/haoran-mc/gcbbs/pkg/config"
	"github.com/haoran-mc/gcbbs/pkg/utils"
)

func Run() {
	engine := gin.Default()

	// 定义在模板中使用的复杂的函数
	engine.SetFuncMap(utils.GetTemplateFuncMap())

	engine.Static("/assets", "../assets")   // 静态文件位置
	engine.LoadHTMLGlob("../views/**/**/*") // 模板位置

	// 使用 cookie 存储 session
	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	// 使用 session.name 作为浏览器中 cookie 的键
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine) // 管理员后台
	route.RegisterFrontedRoute(engine) // 普通用户前台

	if err := engine.Run(":8082"); err != nil {
		log.Fatalf("serevr running error: %v", err)
	}
}
