package user

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/user"
)

func Auth(ctx *gin.Context) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		//return vaild fail
		resource.Unauthorized(ctx, errors.New("沒有token"))
		return
	}

	token := strings.TrimSpace(splitToken[1])

	verifyResult := helper.ValidJSONWebToken(token)
	if !verifyResult {
		resource.Unauthorized(ctx, errors.New("驗證失敗，請重新登入"))
		return
	}

	u, err := user.Auth(&user.UserAuthRequest{
		Token: token,
	})
	if err != nil {
		resource.Unauthorized(ctx, err)
		return
	}

	resource.Success(ctx, u, nil)
}
