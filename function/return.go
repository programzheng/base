package function

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Value interface{}
	Error error
}

func Response(ctx *gin.Context, vaild error, value interface{}, err error) {
	if vaild != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": vaild.Error(),
		})
	} else {
		result := Result{
			Value: value,
			Error: err,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Code":   http.StatusOK,
			"Result": result,
		})
	}
}
