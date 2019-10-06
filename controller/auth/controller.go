package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ProgramZheng/base/function"
	"github.com/gin-gonic/gin"
)

func Vaild(ctx *gin.Context) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		//return not vaild
		ctx.AbortWithStatus(401)
		return
	}

	requestToken = strings.TrimSpace(splitToken[1])

	claims, err := function.ValidJSONWebToken(requestToken)
	if claims == nil {
		err = errors.New("請重新登入")
		function.Response(ctx, nil, false, err)
		ctx.AbortWithStatus(401)
		return
	}
	fmt.Println(claims)
	if err != nil {
		function.Response(ctx, nil, claims, err)
		ctx.AbortWithStatus(200)
		return
	}
	function.Response(ctx, nil, true, nil)
	return
}
