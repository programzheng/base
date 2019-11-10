package auth

import (
	"errors"
	"strings"

	"github.com/ProgramZheng/base/function"
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

	token := strings.TrimSpace(splitToken[1])

	vaildResult := function.ValidJSONWebToken(token)
	if !vaildResult {
		function.Unauthorized(ctx, errors.New("請重新登入2"))
		return
	}

	// authService := auth_service.AdminLogin{
	// 	Token: token,
	// }
	// adminLogin, err := authService.GetAdminLogin()
	// if adminLogin.ID == 0 && err != nil {
	// 	function.Unauthorized(ctx, errors.New("請重新登入"))
	// 	return
	// }

	function.Success(ctx, nil, nil)
	return
}
