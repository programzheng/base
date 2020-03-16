package router

import (
	"github.com/programzheng/base/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/controller/admin"
	"github.com/programzheng/base/pkg/controller/auth"
	"github.com/programzheng/base/pkg/controller/bot"
	"github.com/programzheng/base/pkg/controller/file"
	"github.com/programzheng/base/pkg/controller/job"
	"github.com/programzheng/base/pkg/controller/post"
)

func SetRouter(router *gin.Engine) {
	setMiddleware(router)
	//設置default Route
	setRoute(router)
	//設置API Router
	setAPIRoute(router)
	//設置Bot Router
	setBotRouter(router)
}

func setMiddleware(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
}

func setRoute(router *gin.Engine) {
	router.GET("/jobrunner/json", job.JobJson)

	router.LoadHTMLGlob("dist/view/*")

	router.GET("/jobrunner/html", job.JobHtml)

	router.GET("files/:hash_id", file.Get)
}

func setAPIRoute(router *gin.Engine) {
	apiGroup := router.Group("/API")
	adminGroup := apiGroup.Group("/admins")
	{
		adminGroup.POST("", admin.Register)
		adminGroup.POST("login", admin.Login)
		adminGroup.POST("auth", auth.VaildAdmin)
	}
	apiGroup.Use(middleware.ValidJSONWebToken())
	{
		adminsGroup := apiGroup.Group("/admins")
		{
			adminsGroup.GET("", admin.Get)
		}
		filesGroup := apiGroup.Group("/files")
		{
			filesGroup.POST("", file.Upload)
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
	}

}

func setBotRouter(router *gin.Engine) {
	botGroup := router.Group("/bot")
	lineGroup := botGroup.Group("/line")
	{
		lineGroup.POST("", bot.LineWebHook)
	}
}
