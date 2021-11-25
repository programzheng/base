package router

import (
	"base/pkg/controller/file"
	"base/pkg/controller/job"
	"base/pkg/controller/socketio"
	"base/pkg/filesystem"

	"github.com/gin-gonic/gin"
)

func setRoute(router *gin.Engine) {
	router.GET("/jobrunner/json", job.JobJson)

	// router.LoadHTMLGlob("dist/view/*")

	// router.GET("/jobrunner/html", job.JobHtml)

	router.GET("files/:hash_id", file.Get)

	router.Static("static", filesystem.Driver.GetPath())
}

func setTestRoute(router *gin.Engine) {
	testGroup := router.Group("/test")
	{
		testGroup.GET("socketio", socketio.View)
	}
}
