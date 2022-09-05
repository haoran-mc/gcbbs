package remind

import (
	"log"

	"github.com/haoran-mc/gcbbs/internal/consts"
	"github.com/haoran-mc/gcbbs/internal/model"
)

// FollowObs 关注提醒
type FollowObs struct {
	Sender   uint64
	Receiver uint64
}

// Update 回复评论提醒
func (o *FollowObs) Update() {
	r := model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      0,
		SourceContent: "",
		SourceType:    consts.UserSource,
		SourceUrl:     "",
		Action:        consts.FollowUserRemind,
	})
	if r.Error != nil {
		log.Panicln(r.Error)
	}
}
