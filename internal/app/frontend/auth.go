package frontend

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/service"
	"github.com/haoran-mc/gcbbs/internal/service/frontend"
)

type auth struct{}

var Auth = auth{}

// RegisterPage 注册页面
func (c *auth) RegisterPage(ctx *gin.Context) {
	service.Context(ctx).View("frontend.auth.register", gin.H{})
}

// RegisterSubmit 注册提交
func (c *auth) RegisterSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/") // 读取参数，当参数不存在的时候，提供一个默认值

	var req fe.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil { // TODO 验证还是？
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	// 为什么一定要写成方法？直接写成函数有什么不好的地方？
	// 每次调用方法前还要使用 UserService 函数创建一个 sUser 结构体
	// 是不是 ctx 的关系，需要上下文，直接从上下文中获取到用户信息
	if err := frontend.UserService(ctx).Register(&req); err != nil { // 调用 service 层服务
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithMsg("注册成功，欢迎来到酷社区").Redirect()
	}
}

// LoginPage 登录页面
func (c *auth) LoginPage(ctx *gin.Context) {
	p := ctx.DefaultQuery("back", "/")
	s := service.Context(ctx)

	if s.Check() {
		s.To(p).Redirect()
	} else {
		s.View("frontend.auth.login", gin.H{"path": p})
	}
}

// LoginSubmit 登录提交
func (c *auth) LoginSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/")

	var req fe.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := frontend.UserService(ctx).Login(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithError("登陆成功，欢迎来到酷社区").Redirect()
	}
}

// LogoutSubmit 用户登出
func (c *auth) LogoutSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	frontend.UserService(ctx).Logout()
	s.To("/").WithMsg("退出成功").Redirect()
}
