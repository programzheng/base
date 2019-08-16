package admin

import (
	"net/http"

	"github.com/ProgramZheng/base/model/admin"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	adminStruct := admin.Admin{}
	ctx.BindJSON(&adminStruct)
	result := admin.Add(adminStruct)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func Login(ctx *gin.Context) {
	query := map[string]interface{}{}
	ctx.BindJSON(&query)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": query,
	})
}
