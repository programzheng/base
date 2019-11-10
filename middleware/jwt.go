package middleware

import (
	"errors"
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
			function.Unauthorized(ctx, errors.New("請重新登入"))
			return
		}
		requestToken = strings.TrimSpace(splitToken[1])

		result := function.ValidJSONWebToken(requestToken)
		if result {
			ctx.Next()
		}
	}
}
