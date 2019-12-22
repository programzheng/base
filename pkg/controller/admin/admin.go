package admin

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/service/admin"
	"github.com/programzheng/base/pkg/service/auth"
)

func Register(ctx *gin.Context) {
	adminService := admin.Admin{}
	if err := ctx.Bind(&adminService); err != nil {
		function.BadRequest(ctx, err)
		return
	}

	//hash password
	adminService.Password = function.CreateHash(adminService.Password)
	if err := adminService.Add(); err != nil {
		function.Fail(ctx, err)
		return
	}

	function.Success(ctx, nil, nil)
	return
}

func Login(ctx *gin.Context) {
	login := auth.Login{}
	if err := ctx.Bind(&login); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	admin, err := (&admin.Admin{
		Account: login.Account,
	}).Get()
	if err != nil {
		function.Fail(ctx, errors.New("帳號錯誤"))
		return
	}
	err = function.CheckHash(admin.Password, login.Password)
	if err != nil {
		function.Fail(ctx, errors.New("密碼錯誤"))
		return
	}
	token := function.CreateJWT()
	adminLogin := auth.AdminLogin{
		AdminID: admin.ID,
		Token:   token.Token,
		IP:      ctx.ClientIP(),
	}
	if err := adminLogin.AddAdminLogin(); err != nil {
		function.Fail(ctx, err)
		return
	}

	function.Success(ctx, token, nil)
	return
}

func GetAdmins(ctx *gin.Context) {
	adminService := admin.Admin{}
	if err := ctx.Bind(&adminService); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	admins, err := adminService.GetAdmins()
	if err != nil {
		function.Fail(ctx, err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = admins
	// data["Total"] = total
	function.Success(ctx, data, nil)
	return
}
