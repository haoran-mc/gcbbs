package frontend

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/service"
	"github.com/haoran-mc/gcbbs/internal/service/frontend"
)

type cComment struct{}

var Comment = cComment{}

// PublishSubmit 提交评论
func (c *cComment) PublishSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.SubmitCommentReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	to := fmt.Sprintf("/topics/%d", req.TopicId)

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.To(to).WithError(err.FirstError()).Redirect()
		return
	}

	if id, err := frontend.CommentService(ctx).Submit(&req); err != nil {
		s.To(to).WithError(err).Redirect()
	} else {
		s.To(fmt.Sprintf("%s?j=comment%d", to, id)).WithMsg("发布成功").Redirect()
	}
}

// DeleteSubmit 删除评论
func (c *cComment) DeleteSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.DeleteCommentReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.FirstError()})
		return
	}

	if err := frontend.CommentService(ctx).Delete(req.ID); err != nil {
		s.Json(gin.H{"code": 1, "msg": "删除失败"})
	} else {
		s.Json(gin.H{"code": 0, "msg": "删除成功"})
	}
}
