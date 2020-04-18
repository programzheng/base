package websocket

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/server/websocket/router"
	"github.com/spf13/viper"
)

func Run() error {
	jobrunner.Start()
	api := gin.Default()
	router.SetRouter(api)
	port := viper.Get("WEBSOCKET_PORT")
	if port != nil {
		return api.Run(":" + port.(string))
	}
	return api.Run()
}
