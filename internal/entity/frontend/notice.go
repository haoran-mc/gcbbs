package frontend

import "github.com/haoran-mc/gcbbs/internal/model"

type SystemUserNotice struct {
	model.SystemUserNotices
	Notice model.SystemNotices `gorm:"foreignKey:notice_id"`
}
