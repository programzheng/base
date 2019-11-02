package admin

import (
	"errors"

	"github.com/ProgramZheng/base/function"
	"github.com/ProgramZheng/base/model"
	"github.com/ProgramZheng/base/model/admin"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	adminStruct := admin.Admin{}
	if err := ctx.Bind(&adminStruct); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	//hash password
	adminStruct.Password = function.CreateHash(adminStruct.Password)
	model.Migrate(&adminStruct, &adminStruct.Profile)
	if err := model.Add(&adminStruct); err != nil {
		function.BadRequest(ctx, err)
		return
	}

	function.Success(ctx, adminStruct, nil)
	return
}

func Login(ctx *gin.Context) {
	login := admin.Login{}
	if err := ctx.Bind(&login); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	where := map[string]interface{}{
		"account": login.Account,
	}
	adminStruct, err := model.Get(&admin.Admin{}, where)
	if adminStruct.(*admin.Admin).ID == 0 || err != nil {
		function.Fail(ctx, errors.New("帳號錯誤"))
		return
	}
	err = function.CheckHash(adminStruct.(*admin.Admin).Password, login.Password)
	if err != nil {
		function.Fail(ctx, errors.New("密碼錯誤"))
		return
	}
	token := function.CreateJWT()
	function.Success(ctx, token, nil)
	return
}
