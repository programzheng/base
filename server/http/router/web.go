package router

import (
	"path/filepath"

	"github.com/programzheng/base/pkg/controller/file"
	"github.com/programzheng/base/pkg/controller/job"
	"github.com/programzheng/base/pkg/controller/socketio"
	"github.com/programzheng/base/pkg/helper"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func setRoute(router *gin.Engine) {
	router.GET("/jobrunner/json", job.JobJson)

	// router.LoadHTMLGlob("dist/view/*")

	// router.GET("/jobrunner/html", job.JobHtml)

	router.StaticFile("files/image/empty", filepath.Join(helper.RootPath, "/dist/image/empty.png"))
	router.GET("files/:hash_id", file.Get)

	if viper.Get("STATIC_UPLOAD_ROUTE").(string) == "true" {
		router.Static("static", filepath.Join(helper.RootPath, "/storage/upload"))
	}
}

func setTestRoute(router *gin.Engine) {
	testGroup := router.Group("/test")
	{
		testGroup.GET("socketio", socketio.View)
	}
}
