package router

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	// 中間件
	setMiddleware(router)
	// 設置Web Route
	setRoute(router)
	// 設置API Router
	setAPIRoute(router)
	// 設置Bot Router
	setBotRouter(router)
}
