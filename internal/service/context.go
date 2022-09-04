package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/model"
	"github.com/haoran-mc/gcbbs/pkg/config"
)

const (
	versionKey = "version"
	errKey     = "err"
	msgKey     = "msg"
	dataKey    = "data"
	userKey    = "user"
	unreadKey  = "unread"
)

// BaseContext encapsulated context
type BaseContext struct {
	Ctx     *gin.Context
	session sessions.Session
	path    string
}

// Context initialize Basecontext
func Context(ctx *gin.Context) *BaseContext {
	stx := &BaseContext{
		Ctx:     ctx,
		session: sessions.Default(ctx),
		path:    "/",
	}
	return stx
}

// Redirect
func (c *BaseContext) Redirect() {
	c.Ctx.Redirect(http.StatusFound, c.path)
}

// clear session
func (c *BaseContext) clear() {
	c.session.Delete(errKey)
	c.session.Delete(msgKey)
	_ = c.session.Save() // store the value in session into cookie
}

// Back to previous page
func (c *BaseContext) Back() *BaseContext {
	c.path = c.Ctx.Request.RequestURI
	return c
}

// To Setting Jump Path
func (c *BaseContext) To(to string) *BaseContext {
	c.path = to
	return c
}

// WithError Error Message Jump
func (c *BaseContext) WithError(err interface{}) *BaseContext {
	errStr := ""

	switch v := err.(type) {
	case error:
		errStr = v.Error()
	case string:
		errStr = v
	}

	c.session.Set(errKey, errStr)
	_ = c.session.Save()
	return c
}

// WithMsg 提示消息跳转
func (c *BaseContext) WithMsg(msg string) *BaseContext {
	c.session.Set(msgKey, msg)
	_ = c.session.Save()
	return c
}

// SetAuth 设置授权
func (c *BaseContext) SetAuth(users model.Users) {
	s, _ := json.Marshal(users)
	c.session.Set(userKey, string(s))
	_ = c.session.Save()
}

// Auth 获取授权
func (c *BaseContext) Auth() *model.Users {
	var user *model.Users
	str := c.session.Get(userKey)
	if str != nil {
		if v, ok := str.(string); ok {
			_ = json.Unmarshal([]byte(v), &user)
		}
	}
	return user
}

// Check 检查授权
func (c *BaseContext) Check() bool {
	user := c.Auth()
	if user == nil {
		return false
	} else {
		return user.IsAdmin > 0
	}
}

// Forget 清理授权
func (c *BaseContext) Forget() {
	c.session.Delete(userKey)
	_ = c.session.Save()
}

// unread 消息未读数
func (c *BaseContext) unread() bool {
	if !c.Check() {
		return false
	}
	var (
		remind *model.Reminds
		notice *model.SystemUserNotices
	)

	UID := c.Auth().ID

	// 提醒消息
	r := model.Remind().M.Where("receiver", UID).Where("readed_at is null").Find(&remind)
	if r.Error == nil && r.RowsAffected > 0 {
		return true
	}

	// 未读系统消息
	s := model.SystemUserNotice().M.Where("user_id", UID).Where("readed_at is null").Find(&notice)
	if s.Error == nil && s.RowsAffected > 0 {
		return true
	}

	return false
}

// View 模板返回
func (c *BaseContext) View(tpl string, data interface{}) {
	obj := gin.H{
		versionKey: config.Conf.App.Version,
		errKey:     c.session.Get(errKey),
		msgKey:     c.session.Get(msgKey),
		userKey:    c.Auth(),
		unreadKey:  c.unread(),
		dataKey:    data,
	}

	c.clear()

	c.Ctx.HTML(http.StatusOK, tpl, obj)
}

// Json 通用 JSON 响应
func (c *BaseContext) Json(data interface{}) {
	c.Ctx.JSON(http.StatusOK, data)
}

// MDFileJson markdown上传图片响应
func (c *BaseContext) MDFileJson(ok int, msg, url string) {
	c.Ctx.JSON(http.StatusOK, gin.H{
		"success": ok,
		"message": msg,
		"url":     url,
	})
}
