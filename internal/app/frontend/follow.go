package frontend

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/service"
	"github.com/haoran-mc/gcbbs/internal/service/frontend"
)

type cFollow struct{}

var Follow = cFollow{}

// FollowSubmit 关注提交
func (c *cFollow) FollowSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登陆后再继续操作"})
		return
	}

	var req fe.FollowUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": "参数错误"})
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.FirstError()})
		return
	}

	if state, err := frontend.UserService(ctx).Follow(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok", "state": state})
	}
}
