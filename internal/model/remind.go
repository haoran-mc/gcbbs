package model

import (
	"time"

	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

// Reminds 提醒
type Reminds struct {
	Model
	Sender        uint64     `gorm:"column:sender" db:"sender" json:"sender" form:"sender"`                                 // 发送人id
	Receiver      uint64     `gorm:"column:receiver" db:"receiver" json:"receiver" form:"receiver"`                         // 接受者id
	SourceId      uint64     `gorm:"column:source_id" db:"source_id" json:"source_id" form:"source_id"`                     // 资源id
	SourceType    string     `gorm:"column:source_type" db:"source_type" json:"source_type" form:"source_type"`             // 资源类型
	SourceContent string     `gorm:"column:source_content" db:"source_content" json:"source_content" form:"source_content"` // 资源内容
	SourceUrl     string     `gorm:"column:source_url" db:"source_url" json:"source_url" form:"source_url"`                 // 提醒发生地址
	Action        string     `gorm:"column:action" db:"action" json:"action" form:"action"`                                 // 动作类型
	ReadedAt      *time.Time `gorm:"column:readed_at" db:"readed_at" json:"readed_at" form:"readed_at"`                     // 阅读时间
}

type remindModel struct {
	M     *gorm.DB
	Table string
}

// Remind ...
func Remind() *remindModel {
	return &remindModel{
		M:     db.DB.Model(&Reminds{}),
		Table: "reminds",
	}
}
