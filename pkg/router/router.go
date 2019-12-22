package router

import (
	"github.com/programzheng/base/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/controller/admin"
	"github.com/programzheng/base/pkg/controller/auth"
	"github.com/programzheng/base/pkg/controller/post"
	"github.com/programzheng/base/pkg/controller/upload"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	setMiddleware()
}

func setMiddleware() {
	Router.Use(middleware.CORSMiddleware())
}

func SetRouter() {
	apiGroup := Router.Group("/API")
	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.POST("register", admin.Register)
		adminGroup.POST("login", admin.Login)
		adminGroup.POST("auth", auth.VaildAdmin)
	}
	apiGroup.Use(middleware.ValidJSONWebToken())
	{
		adminsGroup := apiGroup.Group("/admins")
		{
			adminsGroup.GET("", admin.Get)
		}
		uploadGroup := apiGroup.Group("/upload")
		{
			uploadGroup.POST("file", upload.File)
		}
		postsGroup := apiGroup.Group("/posts")
		{
			postsGroup.POST("", post.Add)
			// postsGroup.GET("/:id", post.GetForID)
			postsGroup.GET("", post.Get)
			// postsGroup.PATCH("/:id", post.SaveForID)
			// postsGroup.PATCH("", post.Save)
			// postsGroup.DELETE("/:id", post.DelForID)
			// postsGroup.DELETE("", post.Del)
		}
		// postsGroup := apiGroup.Group("/post")
		// {
		// 	postsGroup.POST("", post.Add)
		// 	postsGroup.GET("/:id", post.GetForID)
		// 	postsGroup.GET("", post.Get)
		// 	postsGroup.PATCH("/:id", post.SaveForID)
		// 	postsGroup.PATCH("", post.Save)
		// 	postsGroup.DELETE("/:id", post.DelForID)
		// 	postsGroup.DELETE("", post.Del)
		// }
	}

}
