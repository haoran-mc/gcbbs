package webserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/pkg/utils"
)

func Run() {
	engine := gin.Default()

	// Define functions to handle complex operations in the template
	engine.SetFuncMap(utils.GetTemplateFuncMap())

	if err := engine.Run(":8082"); err != nil {
		log.Fatalf("serevr running error: %v", err)
	}
}
