package admin

import (
	"errors"

	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/admin"
	"github.com/programzheng/base/pkg/service/auth"

	"github.com/gin-gonic/gin"
)

var adminService admin.Admin

func Register(ctx *gin.Context) {
	if err := ctx.Bind(&adminService); err != nil {
		helper.BadRequest(ctx, err)
		return
	}

	//hash password
	adminService.Password = helper.CreateHash(adminService.Password)
	result, err := adminService.Add()
	if err != nil {
		helper.Fail(ctx, err)
		return
	}

	helper.Success(ctx, result, nil)
	return
}

func Login(ctx *gin.Context) {
	login := auth.Login{}
	if err := ctx.Bind(&login); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	token, err := admin.Login(ctx.ClientIP(), login.Account, login.Password)
	if err != nil {
		helper.Fail(ctx, errors.New("登入失敗"))
		return
	}

	helper.Success(ctx, token, nil)
}

func Get(ctx *gin.Context) {
	adminService := admin.Admin{}
	if err := ctx.Bind(&adminService); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	admins, err := adminService.Get()
	if err != nil {
		helper.Fail(ctx, err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = admins
	data["total"] = len(admins)
	helper.Success(ctx, data, nil)
}
