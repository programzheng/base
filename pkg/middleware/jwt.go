package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/resource"

	"github.com/gin-gonic/gin"
)

type authAdminString string

func ValidJSONWebToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paser Header's Authorization
		requestToken := ctx.GetHeader("Authorization")
		splitToken := strings.Split(requestToken, "Bearer")
		if len(splitToken) != 2 {
			//return not vaild
			resource.Unauthorized(ctx, errors.New("請重新登入"))
			return
		}
		requestToken = strings.TrimSpace(splitToken[1])

		result := helper.ValidJSONWebToken(requestToken)
		if !result {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}

func GraphqlValidJSONWebToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paser Header's Authorization
		requestToken := ctx.GetHeader("Authorization")
		splitToken := strings.Split(requestToken, "Bearer")
		if len(splitToken) == 1 {
			ctx.Next()
			return
		}
		requestToken = strings.TrimSpace(splitToken[1])

		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), authAdminString("auth_admin"), requestToken))

		result := helper.ValidJSONWebToken(requestToken)
		if !result {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}

func CtxValue(ctx context.Context) string {
	raw, _ := ctx.Value(authAdminString("auth_admin")).(string)
	return raw
}
