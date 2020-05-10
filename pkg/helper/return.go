package helper

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
		"Result": result,
	})
}

func BadRequest(ctx *gin.Context, err error) {
	if err != nil {
		customError = err.Error()
	} else {
		customError = nil
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Code":    http.StatusBadRequest,
		"Message": customError,
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
		"Code":    http.StatusUnprocessableEntity,
		"Message": customError,
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
		"Code":    http.StatusUnauthorized,
		"Message": customError,
	})
	ctx.AbortWithStatus(http.StatusUnauthorized)
}

func Success(ctx *gin.Context, value interface{}, message interface{}) {
	if message != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Value":   value,
			"Message": message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Code":  http.StatusOK,
			"Value": value,
		})
	}
	ctx.AbortWithStatus(http.StatusOK)

}

func UploadSuccess(ctx *gin.Context, value interface{}, message interface{}) {
	if message != nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"Code":    http.StatusCreated,
			"Value":   value,
			"Message": message,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"Code":  http.StatusCreated,
			"Value": value,
		})
	}
	ctx.AbortWithStatus(http.StatusCreated)

}
