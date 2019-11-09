package auth

import (
	"errors"
	"strings"

	"github.com/ProgramZheng/base/function"
	"github.com/ProgramZheng/base/model"
	"github.com/ProgramZheng/base/model/admin"
	"github.com/gin-gonic/gin"
)

func VaildAdmin(ctx *gin.Context) {
	requestToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer")
	if len(splitToken) != 2 {
		//return not vaild
		function.Unauthorized(ctx, errors.New("驗證失敗"))
		return
	}

	requestToken = strings.TrimSpace(splitToken[1])

	claims, err := function.ValidJSONWebToken(requestToken)
	if err != nil {
		function.Unauthorized(ctx, errors.New("請重新登入"))
		return
	}

	where := map[string]interface{}{
		"token": requestToken,
	}
	adminLoginStruct, err := model.Get(&admin.AdminLogin{}, where)
	if err != nil {
		function.Unauthorized(ctx, errors.New("請重新登入"))
		return
	}

	function.Success(ctx, claims, err)
	return
}
