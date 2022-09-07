package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/service"
)

type sRemind struct {
	ctx *service.BaseContext
}

// RemindService ...
func RemindService(ctx *gin.Context) *sRemind {
	return &sRemind{
		ctx: service.Context(ctx),
	}
}
