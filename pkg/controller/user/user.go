package user

import (
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/user"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var ur user.UserRequest
	if err := ctx.Bind(&ur); err != nil {
		resource.BadRequest(ctx, err)
		return
	}

	result, err := ur.GenerateUser()
	if err != nil {
		resource.Fail(ctx, err)
		return
	}

	resource.Success(ctx, result, nil)
}

func Login(ctx *gin.Context) {
	var ulgr user.UserLoginRequest
	if err := ctx.Bind(&ulgr); err != nil {
		resource.BadRequest(ctx, err)
		return
	}

	user, err := ulgr.Login()
	if err != nil {
		resource.Fail(ctx, err)
		return
	}

	resource.Success(ctx, user, nil)
}
