package job

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func JobJson(ctx *gin.Context) {
	ctx.JSON(200, jobrunner.StatusJson())
}

func JobHtml(ctx *gin.Context) {
	ctx.HTML(200, "jobrunner.html", jobrunner.StatusPage())
}
