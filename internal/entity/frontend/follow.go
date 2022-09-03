package frontend

import "github.com/haoran-mc/gcbbs/internal/model"

type Follow struct {
	model.Follows
	Follower *model.Users `gorm:"foreignKey:user_id"`
	Fans     *model.Users `gorm:"foreignKey:target_id"`
}
