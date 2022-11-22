package frontend

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/consts"
	"github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/model"
	"github.com/haoran-mc/gcbbs/internal/service"
	remindSub "github.com/haoran-mc/gcbbs/internal/subject/remind"
	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

type sLike struct {
	ctx *service.BaseContext
}

// LikeService ...
func LikeService(ctx *gin.Context) *sLike {
	return &sLike{
		ctx: service.Context(ctx),
	}
}

// IsLiked 是否已点赞
func (s *sLike) IsLiked(id uint64, source string) (bool, error) {
	user := s.ctx.Auth()

	var like *model.Likes
	f := model.Like().M.Where(&model.Likes{
		UserId:     user.ID,
		SourceType: source,
		SourceId:   id,
	}).Find(&like)
	if f.Error != nil {
		return false, f.Error
	} else {
		return like.ID > 0, nil
	}
}

// Like 点赞提交
func (s *sLike) Like(req *frontend.LikeReq) error {
	liked, err := s.IsLiked(req.SourceID, req.SourceType)
	if err != nil {
		return errors.New("点赞失败，请稍后再试")
	}

	if liked {
		return errors.New("无法重复点赞")
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		c := tx.Create(&model.Likes{
			UserId:       s.ctx.Auth().ID,
			SourceType:   req.SourceType,
			SourceId:     req.SourceID,
			TargetUserId: req.TargetUserID,
			State:        consts.Liked,
		})
		if c.Error != nil || c.RowsAffected <= 0 { // TODO 整理所有可能遇到的错误！！
			return errors.New("点赞失败，请稍后再试")
		}

		data := map[string]interface{}{
			"like_count": gorm.Expr("like_count + ?", 1),
		}
		if req.SourceType == consts.TopicSource { // 喜欢话题/文章
			// TODO 一定要 &model.Topics{} 吗
			u := tx.Model(&model.Topics{}).Where("id", req.SourceID).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				return errors.New("点赞失败，请稍候再试")
			}
			return nil
		}

		// 如果不是给话题爱心，那么就是给评论爱心
		u := tx.Model(&model.Comments{}).Where("id", req.SourceID).Updates(data)
		if u.Error != nil || u.RowsAffected <= 0 {
			return errors.New("点赞失败，请稍后再试")
		}

		return nil
	})

	if err != nil {
		return err
	}

	sub := remindSub.New()
	sub.Attach(&remindSub.LikeObs{
		Sender:     s.ctx.Auth().ID,
		Receiver:   req.TargetUserID,
		SourceID:   req.SourceID,
		SourceType: req.SourceType,
	})
	sub.Notify()

	return nil
}
