package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/controller"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/user"
)

func Auth(ctx *gin.Context) {
	token, err := controller.GetTokenByGinContext(ctx)
	if err != nil {
		resource.Unauthorized(ctx, err)
	}

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
