package http

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/server/http/router"
	"github.com/spf13/viper"
)

func Run() error {
	jobrunner.Start()
	route := gin.Default()
	router.SetRouter(route)
	port := viper.Get("APP_PORT")
	if port != nil {
		return route.Run(":" + port.(string))
	}
	return route.Run()
}
