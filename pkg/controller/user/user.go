package user

import (
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/user"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var userRequest user.UserRequest
	if err := ctx.Bind(&userRequest); err != nil {
		resource.BadRequest(ctx, err)
		return
	}

	//hash password
	userRequest.Password = helper.CreateHash(userRequest.Password)
	result, err := userRequest.GenerateUser()
	if err != nil {
		resource.Fail(ctx, err)
		return
	}

	resource.Success(ctx, result, nil)
}
