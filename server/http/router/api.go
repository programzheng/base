package router

import (
	"github.com/gin-gonic/gin"

	"github.com/programzheng/base/pkg/controller/admin"
	"github.com/programzheng/base/pkg/controller/auth"
	"github.com/programzheng/base/pkg/controller/file"
	"github.com/programzheng/base/pkg/controller/post"
	"github.com/programzheng/base/pkg/controller/user"
	"github.com/programzheng/base/pkg/middleware"
	"github.com/programzheng/base/pkg/resource"
)

func setAPIRoute(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")
	adminGroup := apiGroup.Group("/admins")
	{
		adminGroup.POST("", admin.Register)
		adminGroup.POST("login", admin.Login)
		adminGroup.POST("vaild_admin_login_log", auth.VaildAdminLoginLog)
	}
	adminGroup.Use(middleware.ValidJSONWebToken())
	{
		adminGroup.POST("auth", func(ctx *gin.Context) {
			resource.Success(ctx, nil, "success")
		})
	}
	userGroup := apiGroup.Group("/users")
	{
		userGroup.POST("", user.Register)
		userGroup.POST("login", user.Login)
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
			postsGroup.PUT("/:id", post.PutByID)
			// postsGroup.PATCH("/:id", post.SaveForID)
			// postsGroup.PATCH("", post.Save)
			postsGroup.DELETE("/:id", post.DelByID)
			// postsGroup.DELETE("", post.Del)
		}
	}

}
