package router

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	// 中間件
	// setMiddleware(router)
	// 設置WebSocket Route
	setWebSocketRouter(router)
}
