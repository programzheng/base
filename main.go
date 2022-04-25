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
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/programzheng/base/config"

	"github.com/programzheng/base/server/http/router"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// cmd.Execute()
	Run()
}

func Run() {
	jobrunner.Start()
	if env := viper.Get("APP_ENV"); env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()
	router.SetRouter(route)

	port := "80"
	if port = viper.GetString("PORT"); port == "" {
		port = viper.GetString("APP_PORT")
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: route,
	}

	log.Printf("listen port: %s\n", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exit")
}
