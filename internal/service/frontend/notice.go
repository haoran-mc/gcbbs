package frontend

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/haoran-mc/gcbbs/internal/consts"
	fe "github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/model"
	"github.com/haoran-mc/gcbbs/internal/service"
	"github.com/haoran-mc/gcbbs/pkg/utils/page"
	"github.com/jinzhu/gorm"
)

type sNotice struct {
	ctx *service.BaseContext
}

// NoticeService ...
func NoticeService(ctx *gin.Context) *sNotice {
	return &sNotice{
		ctx: service.Context(ctx),
	}
}

// GetList 获取消息列表
func (s *sNotice) GetList(req *fe.GetRemindListReq) (gin.H, error) {
	var (
		total  int64
		limit  = 30
		offset = (req.Page - 1) * limit
	)

	if len(req.Type) <= 0 {
		req.Type = consts.RemindNotice
	}

	if req.Type == consts.RemindNotice { // remind 提醒
		var list []*fe.Remind

		query := model.Remind().M.Where("receiver", s.ctx.Auth().ID)

		if c := query.Count(&total); c.Error != nil {
			return nil, c.Error
		}

		f := query.
			Preload("SenderUser").
			Order("id DESC").
			Limit(limit).
			Offset(offset).
			Find(&list)
		if f.Error != nil && !errors.Is(f.Error, gorm.ErrRecordNotFound) {
			return nil, f.Error
		}

		pageObj := page.New(int(total), limit, gconv.Int(req.Page), s.ctx.Ctx.Request.RequestURI)

		return gin.H{
			"list": list,
			"page": pageObj,
			"type": req.Type,
		}, nil
	} else { // 否则就是系统提醒
		var list []*fe.SystemUserNotice

		query := model.SystemUserNotice().M.Where("user_id", s.ctx.Auth().ID)
		if c := query.Count(&total); c.Error != nil {
			return nil, c.Error
		}

		f := query.
			Preload("Notice").
			Order("id DESC").
			Limit(limit).
			Offset(offset).
			Find(&list)
		if f.Error != nil && !errors.Is(f.Error, gorm.ErrRecordNotFound) {
			return nil, f.Error
		}

		pageObj := page.New(int(total), limit, gconv.Int(req.Page), s.ctx.Ctx.Request.RequestURI)

		return gin.H{
			"list": list,
			"page": pageObj,
			"type": req.Type,
		}, nil
	}
}

// GetRemindUnread 获取提醒未读数
func (s *sNotice) GetRemindUnread() (int64, error) {
	var total int64
	c := model.Remind().M.Where("receiver", s.ctx.Auth().ID).
		Where("readed_at is NULL").Count(&total)
	if c.Error != nil {
		return 0, c.Error
	} else {
		return total, nil
	}
}

// GetLetterUnread TODO 获取私信未读数
func (s *sNotice) GetLetterUnread() (int64, error) {
	return 0, nil
}

// GetSystemUnread() 获取系统未读数
func (s *sNotice) GetSystemUnread() (int64, error) {
	var total int64
	c := model.SystemUserNotice().M.Where("user_id", s.ctx.Auth().ID).
		Where("readed_at is NULL").Count(&total)
	if c.Error != nil {
		return 0, c.Error
	} else {
		return total, nil
	}
}

// ReadAll 读取消息
func (s *sNotice) ReadAll(t string) {
	currUser := s.ctx.Auth()
	if t == consts.RemindNotice {
		model.Remind().M.Where("readed_at is null AND receiver = ?", currUser.ID).
			Update("readed_at", time.Now())
	} else {
		model.SystemUserNotice().M.Where("readed_at is null AND user_id = ?", currUser.ID).
			Update("readed_at", time.Now())
	}
}
