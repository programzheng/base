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
	postGroup := Router.Group("/post")
	{
		postGroup.POST("", post.Add)
		postGroup.GET("/:id", post.GetForID)
		postGroup.GET("", post.Get)
		postGroup.PATCH("/:id", post.SaveForID)
		postGroup.PATCH("", post.Save)
		postGroup.DELETE("/:id", post.DelForID)
		postGroup.DELETE("", post.Del)
	}
}
