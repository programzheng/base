package post

import (
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service"
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
	page := service.Page{}
	if err := ctx.Bind(&page); err != nil {
		helper.BadRequest(ctx, err)
		return
	}

	posts, err := postService.Get(page)
	if err != nil {
		helper.Fail(ctx, err)
		return
	}
	count, err := postService.GetTotalNumber()
	if err != nil {
		helper.Fail(ctx, err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = posts
	data["total"] = count
	data["page"] = page

	helper.Success(ctx, data, nil)
}

func PutByID(ctx *gin.Context) {
	if err := ctx.Bind(&postService); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	id := ctx.Param("id")
	uid := helper.ConvertToUint(id)

	post, err := postService.UpdateByID(uid)
	if err != nil {
		helper.Fail(ctx, err)
		return
	}

	data := map[string]interface{}{
		"post": post,
	}
	helper.Success(ctx, data, nil)
}

func DelByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := helper.ConvertToUint(id)

	err := postService.DelByID(uid)

	if err != nil {
		helper.Fail(ctx, err)
		return
	}

	helper.Success(ctx, nil, "Delete Success")
}
