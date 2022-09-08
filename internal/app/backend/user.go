package backend

import (
	"github.com/gin-gonic/gin"

	be "github.com/haoran-mc/gcbbs/internal/entity/backend"
	sv "github.com/haoran-mc/gcbbs/internal/service"
	bs "github.com/haoran-mc/gcbbs/internal/service/backend"
)

type cUser struct{}

var User = cUser{}

// IndexPage 用户主页
func (c *cUser) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if data, err := bs.UserService(ctx).GetList(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.View("backend.user.index", data)
	}
}
