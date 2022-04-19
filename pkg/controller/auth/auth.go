package auth

import (
	"errors"
	"strings"

	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/auth"

	"github.com/gin-gonic/gin"
)

func VaildAdminLoginLog(ctx *gin.Context) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		//return not vaild
		resource.Unauthorized(ctx, errors.New("驗證失敗"))
		return
	}

	token := strings.TrimSpace(splitToken[1])

	//曾經有登入記錄
	adminLogin, err := (&auth.AdminLogin{
		Token: token,
	}).GetAdminLogin()

	if err != nil {
		resource.Unauthorized(ctx, errors.New("請重新登入"))
		return
	}

	vaildResult := helper.ValidJSONWebToken(token)
	if !vaildResult {
		resource.Unauthorized(ctx, errors.New("請重新登入"))
		return
	}

	if adminLogin.Remember {

	}

	resource.Success(ctx, nil, nil)
	return
}
