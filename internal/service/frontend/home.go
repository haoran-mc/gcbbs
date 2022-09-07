package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/service"
)

type sHome struct {
	ctx *service.BaseContext
}

// HomeService ...
func HomeService(ctx *gin.Context) *sHome {
	return &sHome{
		ctx: service.Context(ctx),
	}
}
