package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Value interface{}
	Error interface{}
}

var customError interface{}

func Response(ctx *gin.Context, value interface{}, err error) {
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
		"result": result,
	})
}

func BadRequest(ctx *gin.Context, err error) {
	if err != nil {
		customError = err.Error()
	} else {
		customError = nil
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": customError,
	})
	ctx.AbortWithStatus(http.StatusBadRequest)
}

func Fail(ctx *gin.Context, err error) {
	if err != nil {
		customError = err.Error()
	} else {
		customError = nil
	}
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"message": customError,
	})
	ctx.AbortWithStatus(http.StatusUnprocessableEntity)
}

func Unauthorized(ctx *gin.Context, err error) {
	if err != nil {
		customError = err.Error()
	} else {
		customError = nil
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": customError,
	})
	ctx.AbortWithStatus(http.StatusUnauthorized)
}

func Success(ctx *gin.Context, value interface{}, message interface{}) {
	if message != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"value":   value,
			"message": message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"value": value,
		})
	}
	ctx.AbortWithStatus(http.StatusOK)

}

func UploadSuccess(ctx *gin.Context, value interface{}, message interface{}) {
	if message != nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusCreated,
			"value":   value,
			"message": message,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"code":  http.StatusCreated,
			"value": value,
		})
	}
	ctx.AbortWithStatus(http.StatusCreated)

}
