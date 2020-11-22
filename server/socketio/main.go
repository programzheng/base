/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	_ "github.com/programzheng/base/config"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/module/socketio"
	"github.com/programzheng/base/server/socketio/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// cmd.Execute()
	Run()
}

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
