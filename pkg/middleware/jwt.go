package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/helper"
)

func ValidJSONWebToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paser Header's Authorization
		requestToken := ctx.GetHeader("Authorization")
		splitToken := strings.Split(requestToken, "Bearer")
		if len(splitToken) != 2 {
			//return not vaild
			helper.Unauthorized(ctx, errors.New("請重新登入"))
			return
		}
		requestToken = strings.TrimSpace(splitToken[1])

		result := helper.ValidJSONWebToken(requestToken)
		if result {
			ctx.Next()
		}
	}
}
