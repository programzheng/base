package router

import (
	"base/pkg/controller/file"
	"base/pkg/controller/job"
	"base/pkg/controller/socketio"

	"github.com/gin-gonic/gin"
)

func setRoute(router *gin.Engine) {
	router.GET("/jobrunner/json", job.JobJson)

	// router.LoadHTMLGlob("dist/view/*")

	// router.GET("/jobrunner/html", job.JobHtml)

	router.GET("files/:hash_id", file.Get)

}

func setTestRoute(router *gin.Engine) {
	testGroup := router.Group("/test")
	{
		testGroup.GET("socketio", socketio.View)
	}
}
