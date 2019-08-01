package router

import (
	"github.com/ProgramZheng/base/controller/post"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.GET("/post/:id", post.Get)
	Router.POST("/post", post.Add)
}
