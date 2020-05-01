package socketio

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func View(ctx *gin.Context) {
	ctx.HTML(200, "socketio.html", gin.H{
		"API": viper.Get("SOCKETIO_URL"),
	})
}
