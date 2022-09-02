package model

import (
	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

// Reports 举报
type Reports struct {
	Model
	UserId     uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                 // 举报人id
	Remark     string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"`                     // 举报备注
	TargetId   uint64 `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"`         // 被举报人id
	SourceId   uint64 `gorm:"column:source_id" db:"source_id" json:"source_id" form:"source_id"`         // 目标id
	SourceType string `gorm:"column:source_type" db:"source_type" json:"source_type" form:"source_type"` // 目标类型
	SourceUrl  string `gorm:"column:source_url" db:"source_url" json:"source_url" form:"source_url"`     // 目标链接
	HandlerId  uint64 `gorm:"column:handler_id" db:"handler_id" json:"handler_id" form:"handler_id"`     // 处理人id
	State      uint8  `gorm:"column:state" db:"state" json:"state" form:"state"`                         // 状态：0-待处理/1-已处理
}

type reportModel struct {
	M     *gorm.DB
	Table string
}

// Report ...
func Report() *reportModel {
	return &reportModel{
		M:     db.DB.Model(&Reminds{}),
		Table: "reports",
	}
}
