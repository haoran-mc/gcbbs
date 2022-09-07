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

	// Define functions to handle complex operations in the template
	engine.SetFuncMap(utils.GetTemplateFuncMap())

	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../views/**/**/*")

	// Use cookies to store sessions
	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	// The session name becomes the key of the cookie in Browser
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	if err := engine.Run(":8082"); err != nil {
		log.Fatalf("serevr running error: %v", err)
	}
}
