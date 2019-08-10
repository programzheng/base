package post

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ProgramZheng/base/model/post"
	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	var postStruct = post.Post{}
	ctx.BindJSON(&postStruct)
	result := post.Add(postStruct)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func GetForID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	result := post.GetForID(post.Post{}, id)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func Get(ctx *gin.Context) {
	where := map[string]interface{}{}
	ctx.BindUri(&where)
	fmt.Println(where)
	result := post.Get(post.Post{}, where)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func SaveForID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	update := map[string]interface{}{}
	ctx.BindJSON(&update)
	result := post.SaveForID(post.Post{}, id, update)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}

func DelForID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	result := post.DelForID(post.Post{}, id)

	ctx.JSON(http.StatusOK, gin.H{
		"Code":   http.StatusOK,
		"Result": result,
	})
}
