package webserver

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	if err := engine.Run(":8082"); err != nil {
		log.Fatalf("serevr running error: %v", err)
	}
}
