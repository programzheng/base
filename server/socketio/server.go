package socketio

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/module/socketio"
	"github.com/programzheng/base/server/socketio/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() error {

	if env := viper.Get("APP_ENV"); env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	// initial socket module
	if err := socketio.NewEngine(); err != nil {
		log.Fatalf("Failed to initialize Socket IO engine: %v", err)
	}
	route := gin.Default()
	router.SetSocketIORoute(route)
	port := viper.Get("SOCKETIO_PORT")
	if port != nil {
		return route.Run(":" + port.(string))
	}
	return route.Run()
}
