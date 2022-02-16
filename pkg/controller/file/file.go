package file

import (
	"base/pkg/helper"
	"base/pkg/service/file"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	hashID := ctx.Param("hash_id")

	files, err := file.Get(nil, func() map[string]interface{} {
		maps := make(map[string]interface{})
		maps["hash_id"] = hashID
		return maps
	})

	if err != nil {
		log.Fatalf("files get error: %v", err)
	}
	if len(files) == 0 {
		helper.BadRequest(ctx, errors.New("files get not found"))
		return
	}
	file := files[len(files)-1]
	img, err := ioutil.ReadFile(filesystem.Create("local").GetLink(file))
	if err != nil {
		log.Fatal(err)
	}
	ctx.Data(http.StatusOK, file.Type, img)
}
