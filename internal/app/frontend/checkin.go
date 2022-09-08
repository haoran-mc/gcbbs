package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/service"
	"github.com/haoran-mc/gcbbs/internal/service/frontend"
)

type cCheckin struct{}

var Checkin = cCheckin{}

// StoreSubmit 提交签到
func (c *cCheckin) StoreSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登陆后再继续操作"})
		return
	}

	if err := frontend.CheckinService(ctx).Store(); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok"})
	}
}
