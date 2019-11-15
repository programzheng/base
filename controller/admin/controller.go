package admin

import (
	"errors"

	"github.com/ProgramZheng/base/function"
	"github.com/ProgramZheng/base/service/admin_service"
	"github.com/ProgramZheng/base/service/auth_service"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	adminService := admin_service.Admin{}
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
	login := auth_service.Login{}
	if err := ctx.Bind(&login); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	adminService := admin_service.Admin{
		Account: login.Account,
	}
	admin, err := adminService.Get()
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
	adminLogin := auth_service.AdminLogin{
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
	adminService := admin_service.Admin{}
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
	data["List"] = admins
	// data["Total"] = total
	function.Success(ctx, data, nil)
	return
}
