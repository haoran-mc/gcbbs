package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/gcbbs/internal/consts"
	"github.com/haoran-mc/gcbbs/internal/model"
	"github.com/haoran-mc/gcbbs/internal/service"
)

type sNode struct {
	ctx *service.BaseContext
}

// NodeService ...
func NodeService(ctx *gin.Context) *sNode {
	return &sNode{
		ctx: service.Context(ctx),
	}
}

// GetEnableNodes 获取已开启的所有节点
func (s *sNode) GetEnableNodes() ([]*model.Nodes, error) {
	var nodes []*model.Nodes
	r := model.Node().M.Where("state", consts.EnableState).Find(&nodes)
	if r.Error != nil {
		return nil, r.Error
	} else {
		return nodes, nil
	}
}
