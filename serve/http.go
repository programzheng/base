package serve

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/config"
	"github.com/programzheng/base/server/http/router"
)

func RunHttpServer() {
	jobrunner.Start()
	if config.GetProductionStatus() {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()
	router.SetRouter(route)

	port := "80"
	if port = config.Cfg.GetString("PORT"); port == "" {
		port = config.Cfg.GetString("APP_PORT")
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
