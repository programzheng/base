package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/middleware"
)

func setMiddleware(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
}
