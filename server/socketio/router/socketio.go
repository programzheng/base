package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/module/socketio"
)

func SetSocketIORoute(router *gin.Engine) {
	router.GET("/socket.io/", socketio.Handler())
	router.POST("/socket.io/", socketio.Handler())
	router.Handle("WS", "/socket.io", socketio.Handler())
	router.Handle("WSS", "/socket.io", socketio.Handler())
}
