package router

import (
	"github.com/programzheng/base/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/controller/admin"
	"github.com/programzheng/base/pkg/controller/auth"
	"github.com/programzheng/base/pkg/controller/bot"
	"github.com/programzheng/base/pkg/controller/file"
	"github.com/programzheng/base/pkg/controller/post"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	setMiddleware()
	//設置API Router
	setAPIRouter()
	//設置Bot Router
	setBotRouter()
}

func setMiddleware() {
	Router.Use(middleware.CORSMiddleware())
}

func setAPIRouter() {
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
		fileGroup := apiGroup.Group("/files")
		{
			fileGroup.POST("", file.Upload)
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

func setBotRouter() {
	botGroup := Router.Group("/bot")
	lineGroup := botGroup.Group("/line")
	{
		lineGroup.POST("", bot.LineWebHook)
	}
}
