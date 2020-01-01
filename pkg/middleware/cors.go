package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
	// return func(ctx *gin.Context) {
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	// ctx.Writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	// 	if ctx.Request.Method == "OPTIONS" {
	// 		ctx.AbortWithStatus(204)
	// 		return
	// 	} else {
	// 		ctx.Next()
	// 	}
	// }
}
