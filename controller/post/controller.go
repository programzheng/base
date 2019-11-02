package post

// import (
// 	"log"
// 	"strconv"

// 	"github.com/ProgramZheng/base/function"
// 	"github.com/ProgramZheng/base/model/post"
// 	"github.com/gin-gonic/gin"
// )

// func Add(ctx *gin.Context) {
// 	var postStruct = post.Post{}
// 	vaild := ctx.BindJSON(&postStruct)
// 	value, err := post.Add(postStruct)

// 	function.Response(ctx, vaild, value, err)
// }

// func GetForID(ctx *gin.Context) {
// 	id, vaild := strconv.Atoi(ctx.Param("id"))
// 	if vaild != nil {
// 		log.Fatal(vaild)
// 	}
// 	value, err := post.GetForID(post.Post{}, id)

// 	function.Response(ctx, vaild, value, err)
// }

// func Get(ctx *gin.Context) {
// 	where := map[string]interface{}{}
// 	vaild := ctx.BindUri(&where)
// 	value, err := post.Get(post.Post{}, where)

// 	function.Response(ctx, vaild, value, err)
// }

// func SaveForID(ctx *gin.Context) {
// 	id, vaild := strconv.Atoi(ctx.Param("id"))
// 	if vaild != nil {
// 		log.Fatal(vaild)
// 	}
// 	update := map[string]interface{}{}
// 	vaild = ctx.BindJSON(&update)
// 	value, err := post.SaveForID(post.Post{}, id, update)

// 	function.Response(ctx, vaild, value, err)
// }

// func Save(ctx *gin.Context) {
// 	params := map[string]interface{}{}
// 	vaild := ctx.BindJSON(&params)
// 	ids, ok := params["ids"]
// 	if ok {
// 		delete(params, "ids")
// 	}
// 	update := params
// 	value, err := post.Save(post.Post{}, ids, update)

// 	function.Response(ctx, vaild, value, err)
// }

// func DelForID(ctx *gin.Context) {
// 	id, vaild := strconv.Atoi(ctx.Param("id"))
// 	if vaild != nil {
// 		log.Fatal(vaild)
// 	}
// 	value, err := post.DelForID(post.Post{}, id)

// 	function.Response(ctx, vaild, value, err)
// }

// func Del(ctx *gin.Context) {
// 	params := map[string]interface{}{}
// 	vaild := ctx.BindJSON(&params)
// 	ids := params["ids"]
// 	value, err := post.Del(post.Post{}, ids)

// 	function.Response(ctx, vaild, value, err)
// }
