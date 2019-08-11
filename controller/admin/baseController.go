package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	query := map[string]interface{}{}
	ctx.BindJSON(&query)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": query,
	})
}
