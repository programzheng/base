package auth

import (
	"github.com/ProgramZheng/base/function"
	"github.com/gin-gonic/gin"
)

func Vaild(ctx *gin.Context) {
	function.Response(ctx, nil, true, nil)
}
