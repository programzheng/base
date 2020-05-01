package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/programzheng/base/pkg/function"
)

func setWebSocketRouter(router *gin.Engine) {
	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	websocketGroup := router.Group("/websocket")
	{
		websocketGroup.GET("echo/*EIO", func(ctx *gin.Context) {
			function.GetJSON(ctx)
			c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
			if err != nil {
				log.Println("upgrade:", err)
				return
			}
			defer func() {
				log.Println("disconnect !!")
				c.Close()
			}()
			for {
				mtype, msg, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("receive: %s\n", msg)
				err = c.WriteMessage(mtype, msg)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		})
	}
}
