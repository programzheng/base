package post

import (
	"net/http"
	"time"

	"github.com/ProgramZheng/base/model/post"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	var postStruct = post.Post{}
	ctx.BindJSON(&postStruct)
	result := post.Get(postStruct)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  result,
		"message": err,
	})
}

func Add(ctx *gin.Context) {
	var postStruct = post.Post{
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	ctx.BindJSON(&postStruct)
	result, err := post.Add(postStruct)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  result,
		"message": err,
	})
}
