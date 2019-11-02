package auth

import (
	"errors"
	"strings"

	"github.com/ProgramZheng/base/function"
	"github.com/gin-gonic/gin"
)

func Vaild(ctx *gin.Context) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		//return not vaild
		function.Unauthorized(ctx, errors.New("驗證失敗"))
		return
	}

	requestToken = strings.TrimSpace(splitToken[1])

	claims, err := function.ValidJSONWebToken(requestToken)
	if claims == nil {
		function.Unauthorized(ctx, errors.New("請重新登入"))
		return
	}
	if err != nil {
		function.Success(ctx, claims, err)
		return
	}
}
