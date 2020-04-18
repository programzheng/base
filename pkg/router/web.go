package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/controller/file"
	"github.com/programzheng/base/pkg/controller/job"
)

func setRoute(router *gin.Engine) {
	router.GET("/jobrunner/json", job.JobJson)

	router.LoadHTMLGlob("dist/view/*")

	router.GET("/jobrunner/html", job.JobHtml)

	router.GET("files/:hash_id", file.Get)
}
