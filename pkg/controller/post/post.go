package post

import (
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/post"

	"github.com/gin-gonic/gin"
)

var postService post.Post

func Add(ctx *gin.Context) {
	if err := ctx.Bind(&postService); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	result, err := postService.Add()
	if err != nil {
		helper.Fail(ctx, err)
		return
	}

	helper.Success(ctx, result, nil)
	return
}

func Get(ctx *gin.Context) {
	if err := ctx.Bind(&postService); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	posts, err := postService.Get()
	if err != nil {
		helper.Fail(ctx, err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = posts
	// data["Total"] = total
	helper.Success(ctx, data, nil)
	return
}
