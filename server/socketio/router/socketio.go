package router

import (
	"base/pkg/module/socketio"

	"github.com/gin-gonic/gin"
)

func SetSocketIORoute(router *gin.Engine) {
	router.GET("/socket.io/", socketio.Handler())
	router.POST("/socket.io/", socketio.Handler())
	router.Handle("WS", "/socket.io", socketio.Handler())
	router.Handle("WSS", "/socket.io", socketio.Handler())
}
