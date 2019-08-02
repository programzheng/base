package post

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ProgramZheng/base/model/post"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	var postStruct = post.Post{
		ID: id,
	}
	result := post.Get(postStruct)
	fmt.Println(result)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func Add(ctx *gin.Context) {
	var postStruct = post.Post{
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	ctx.BindJSON(&postStruct)
	result := post.Add(postStruct)
	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}
