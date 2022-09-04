package frontend

import "github.com/haoran-mc/gcbbs/internal/model"

type GetRemindListReq struct {
	Page int    `form:"page"`
	Type string `form:"type"`
}

type Remind struct {
	model.Reminds
	ReceiverUser model.Users `gorm:"foreignKey:receiver"`
	SenderUser   model.Users `gorm:"foreignKey:sender"`
}