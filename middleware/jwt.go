package middleware

import (
	"strings"

	"github.com/ProgramZheng/base/function"
	"github.com/gin-gonic/gin"
)

func ValidJSONWebToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paser Header's Authorization
		requestToken := ctx.GetHeader("Authorization")
		splitToken := strings.Split(requestToken, "Bearer")
		if len(splitToken) != 2 {
			//return not vaild
			ctx.AbortWithStatus(401)
			return
		}
		requestToken = strings.TrimSpace(splitToken[1])

		claims, err := function.ValidJSONWebToken(requestToken)
		if err != nil {
			function.Response(ctx, nil, claims, err)
			ctx.AbortWithStatus(200)
			return
		}

		ctx.Next()
	}
}
