package backend

import (
	"github.com/gin-gonic/gin"
	srv "github.com/haoran-mc/gcbbs/internal/service"
)

type cHome struct{}

var Home = cHome{}

func (c *cHome) IndexPage(ctx *gin.Context) {
	srv.Context(ctx).View("backend.home.index", gin.H{})
}
