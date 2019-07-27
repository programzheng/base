package post

import (
	"time"

	"github.com/ProgramZheng/base/model/post"
	"github.com/gin-gonic/gin"
)

func DoAdd(ctx *gin.Context) {
	var postStruct = post.Post{
		// Title:      "testqwe",
		// Text:       "I am textasd",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	ctx.BindJSON(&postStruct)
	post.Add(postStruct)
}
