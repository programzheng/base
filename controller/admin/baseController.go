package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	fmt.Println(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,
	})
}
