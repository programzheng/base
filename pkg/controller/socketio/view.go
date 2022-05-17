package socketio

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/config"
)

func View(ctx *gin.Context) {
	ctx.HTML(200, "socketio.html", gin.H{
		"API": config.Cfg.GetString("SOCKETIO_URL"),
	})
}
