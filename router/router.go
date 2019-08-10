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
	Router.POST("/post", post.Add)
	Router.GET("/post/:id", post.GetForID)
	Router.GET("/post", post.Get)
	Router.PATCH("/post/:id", post.SaveForID)
	Router.DELETE("post/:id", post.DelForID)
}
