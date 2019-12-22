package post

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/service/post"
)

var postService post.Post

func Add(ctx *gin.Context) {
	if err := ctx.Bind(&postService); err != nil {
		function.BadRequest(ctx, err)
		return
	}

	if err := postService.Add(); err != nil {
		function.Fail(ctx, err)
		return
	}

	function.Success(ctx, nil, nil)
	return
}

func Get(ctx *gin.Context) {
	if err := ctx.Bind(&postService); err != nil {
		function.BadRequest(ctx, err)
		return
	}
	posts, err := postService.Get()
	if err != nil {
		function.Fail(ctx, err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = posts
	// data["Total"] = total
	function.Success(ctx, data, nil)
	return
}
