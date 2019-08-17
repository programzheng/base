package admin

import (
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
	check := function.CheckHash(admin.Password, login.Password)
	if check == true {
		token := function.CreateJWT()
		fmt.Println(token)
	}
	value := "登入失敗"
	function.Response(ctx, vaild, value, err)
}
