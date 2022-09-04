package model

import (
	"time"

	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

// Checkins 签到
type Checkins struct {
	NoDeleteModel
	UserId         uint64    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                                 // 用户 ID
	CumulativeDays uint64    `gorm:"column:cumulative_days" db:"cumulative_days" json:"cumulative_days" form:"cumulative_days"` // 累积签到（天）
	ContinuityDays uint64    `gorm:"column:continuity_days" db:"continuity_days" json:"continuity_days" form:"continuity_days"` // 连续签到（天）
	LastTime       time.Time `gorm:"column:last_time" db:"last_time" json:"last_time" form:"last_time"`                         // 最后签到时间
}

type checkinModel struct {
	M     *gorm.DB
	Table string
}

func Checkin() *checkinModel {
	return &checkinModel{
		M:     db.DB.Model(&Checkins{}),
		Table: "checkins",
	}
}
