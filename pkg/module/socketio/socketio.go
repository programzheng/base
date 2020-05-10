package socketio

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
)

var SocketIOServer *socketio.Server
var err error

func NewEngine() error {
	SocketIOServer, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	SocketIOServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})
	SocketIOServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	SocketIOServer.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	SocketIOServer.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	SocketIOServer.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})
	SocketIOServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	return nil
}

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Origin", origin)
		SocketIOServer.ServeHTTP(ctx.Writer, ctx.Request)
		go SocketIOServer.Serve()
	}
}
