package model

import (
	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

type SystemNotices struct {
	Model
	UserId    uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             // 发布人id
	Titile    string `gorm:"column:titile" db:"titile" json:"titile" form:"titile"`                 // 消息标题
	TargetId  string `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"`     // 接收者id
	Content   string `gorm:"column:content" db:"content" json:"content" form:"content"`             // 消息内容
	MDContent string `gorm:"column:md_content" db:"md_content" json:"md_content" form:"md_content"` // markdown内容
}

type SystemNoticeModel struct {
	M     *gorm.DB
	Table string
}

// SystemNotice ...
func SystemNotice() *SystemNoticeModel {
	return &SystemNoticeModel{
		M:     db.DB.Model(&SystemNoticeModel{}),
		Table: "system_notices",
	}
}
