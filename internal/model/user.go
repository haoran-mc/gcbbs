package model

import (
	"time"

	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

type Users struct {
	Model
	Name        string     `gorm:"column:name" db:"name" json:"name" form:"name"`                                     // 用户名
	Gender      uint8      `gorm:"column:gender" db:"gender" json:"gender" form:"gender"`                             // 性别：1-男/2-女/0-未知
	City        string     `gorm:"column:city" db:"city" json:"city" form:"city"`                                     // 城市
	Email       string     `gorm:"column:email" db:"email" json:"email" form:"email"`                                 // 用户邮箱
	Avatar      string     `gorm:"column:avatar" db:"avatar" json:"avatar" form:"avatar"`                             // 用户头像
	Integral    uint64     `gorm:"column:integral" db:"integral" json:"integral" form:"integral"`                     // 个人积分
	Site        string     `gorm:"column:site" db:"site" json:"site" form:"site"`                                     // 个人网站
	Job         string     `gorm:"column:job" db:"job" json:"job" form:"job"`                                         // 职业
	Desc        string     `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`                                     // 简介
	Password    string     `gorm:"column:password" db:"password" json:"password" form:"password"`                     // 密码
	IsAdmin     uint8      `gorm:"column:is_admin" db:"is_admin" json:"is_admin" form:"is_admin"`                     // 是否是管理员：1-是/0-否
	State       uint8      `gorm:"column:state" db:"state" json:"state" form:"state"`                                 // 状态：1-正常/0-禁用
	LastLoginIp string     `gorm:"column:last_login_ip" db:"last_login_ip" json:"last_login_ip" form:"last_login_ip"` // 最后登陆IP
	LastLoginAt *time.Time `gorm:"column:last_login_at" db:"last_login_at" json:"last_login_at" form:"last_login_at"` // 最后登陆时间
}

type userModel struct {
	M     *gorm.DB
	Table string
}

// User ...
func User() *userModel {
	return &userModel{
		M:     db.DB.Model(&Users{}),
		Table: "users",
	}
}
