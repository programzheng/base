package function

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Value interface{}
	Error interface{}
}

func Response(ctx *gin.Context, vaild error, value interface{}, err error) {
	if vaild != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": vaild.Error(),
		})
	} else {
		var customError interface{}
		if err != nil {
			customError = err.Error()
		} else {
			customError = nil
		}
		result := Result{
			Value: value,
			Error: customError,
		}
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"Code":   http.StatusOK,
			"Result": result,
		})
	}
}
