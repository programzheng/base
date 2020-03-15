package server

import (
	"github.com/bamzi/jobrunner"
	"github.com/programzheng/base/pkg/router"
	"github.com/spf13/viper"
)

func Run() {
	jobrunner.Start()
	r := router.Router
	port := viper.Get("APP_PORT")
	if port != nil {
		r.Run(":" + port.(string))
	} else {
		r.Run()
	}
}
