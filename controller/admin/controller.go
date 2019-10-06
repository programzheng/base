package admin

import (
	"errors"
	"fmt"

	"github.com/ProgramZheng/base/function"
	"github.com/ProgramZheng/base/model/admin"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	adminStruct := admin.Admin{}
	vaild := ctx.BindJSON(&adminStruct)
	//hash password
	adminStruct.Password = function.CreateHash(adminStruct.Password)
	value, err := admin.Add(adminStruct)

	function.Response(ctx, vaild, value, err)
}

func Login(ctx *gin.Context) {
	login := admin.Login{}
	vaild := ctx.BindJSON(&login)

	where := map[string]interface{}{
		"account": login.Account,
	}
	admin, err := admin.Get(admin.Admin{}, where)
	fmt.Println(admin.ID)
	if admin.ID == 0 {
		err = errors.New("帳號錯誤")
	}
	if err != nil {
		function.Response(ctx, vaild, nil, err)
		return
	}
	err = function.CheckHash(admin.Password, login.Password)
	if err != nil {
		err = errors.New("密碼錯誤")
		function.Response(ctx, vaild, nil, err)
		return
	}
	token := function.CreateJWT()
	function.Response(ctx, vaild, token, err)
	return
}
