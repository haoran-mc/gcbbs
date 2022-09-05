package frontend

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/consts"
	"github.com/haoran-mc/gcbbs/internal/entity/frontend"
	"github.com/haoran-mc/gcbbs/internal/model"
	"github.com/haoran-mc/gcbbs/internal/service"
	remindSub "github.com/haoran-mc/gcbbs/internal/subject/remind"
	"github.com/haoran-mc/gcbbs/pkg/db"
	"gorm.io/gorm"
)

type sComment struct {
	ctx *service.BaseContext
}

// CommentService ...
func CommentService(ctx *gin.Context) *sComment {
	return &sComment{
		ctx: service.Context(ctx),
	}
}

// Submit 提交评论
func (s *sComment) Submit(req *frontend.SubmitCommentReq) (uint64, error) {
	comment := &model.Comments{
		TopicId:   req.TopicId,
		ReplyId:   req.ReplyId,
		TargetId:  req.TargetId,
		UserId:    s.ctx.Auth().ID,
		Content:   req.Content,
		MDContent: req.MDContent,
	}
	r := model.Comment().M.Create(comment)
	if r.Error != nil {
		return 0, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return 0, errors.New("提交失败，请稍后再试")
	}

	data := map[string]interface{}{
		"reply_id":      s.ctx.Auth().ID,
		"comment_count": gorm.Expr("comment_count + ?", 1),
		"last_reply_at": time.Now(),
	}

	r = model.Topic().M.Where("id = ?", req.TopicId).Updates(data)
	if r.Error != nil {
		return 0, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return 0, errors.New("提交失败，请稍后再试")
	}

	if req.ReplyId <= 0 { // TODO 不是回复
		sub := remindSub.New()
		sub.Attach(&remindSub.CommentObs{
			TopicID:   req.TargetId,
			Sender:    s.ctx.Auth().ID,
			CommentId: comment.ID,
		})
		sub.Notify()
	} else { // 是回复
		sub := remindSub.New()
		sub.Attach(&remindSub.ReplyObs{
			TopicID:   req.TopicId,
			Sender:    s.ctx.Auth().ID,
			CommentId: comment.ID,
			Receiver:  req.ReplyId,
		})
		sub.Notify()
	}

	return comment.ID, nil
}

// GetList 获取评论列表
func (s *sComment) GetList(topicId uint64) ([]*frontend.Comment, error) {
	var list []*frontend.Comment

	query := model.Comment().M // *gorm.DB
	if s.ctx.Check() {
		query = query.Preload(
			"Like",
			"user_id = ? AND source_type = ?",
			s.ctx.Auth().ID,
			consts.CommentSource,
		)
	}

	r := query.
		Where("topic_id = ?", topicId).
		Order("id ASC").
		Preload("Publisher").
		Preload("Topic").
		Preload("Responder").
		Find(&list)
	if r.Error != nil {
		return nil, r.Error
	}

	return list, nil
}

// Delete 删除评论
func (s *sComment) Delete(id uint64) error {
	if !s.ctx.Check() {
		return errors.New("权限不足")
	}

	var comment *model.Comments
	f := model.Comment().M.Where("id", id).Find(&comment)
	if f.Error != nil || comment == nil {
		log.Panicln(f.Error)
		return errors.New("删除失败")
	}

	if comment.UserId != s.ctx.Auth().ID {
		return errors.New("权限不足")
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		d := tx.Delete(&model.Comments{}, id)
		if d.Error != nil || d.RowsAffected <= 0 {
			// TODO 为什么有的时候用中文，有的时候用英文，为什么这里用 Errorf
			return fmt.Errorf("delete comment error: %v", d.Error)
		}
		u := tx.Model(&model.Topics{}).Where("id", comment.TopicId).Updates(map[string]interface{}{
			"comment_count": gorm.Expr("comment_count - 1"),
		})
		if u.Error != nil || u.RowsAffected <= 0 {
			return fmt.Errorf("delete comment error: %v", d.Error)
		}
		return nil
	})

	return err
}
