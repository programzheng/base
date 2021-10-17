package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	// 中間件
	setMiddleware(router)
	// 設置Web Route
	setRoute(router)
	fmt.Println(gin.Mode())
	if mode := gin.Mode(); mode != gin.ReleaseMode {
		// 測試頁面 Route
		setTestRoute(router)
	}
	// 設置API Router
	setAPIRoute(router)
}
