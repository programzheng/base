package file

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/file"

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
		resource.BadRequest(ctx, errors.New("files get not found"))
		return
	}
	file := files[len(files)-1]
	img, err := ioutil.ReadFile(file.GetLocalLink())
	if err != nil {
		log.Fatal(err)
	}
	ctx.Data(http.StatusOK, file.Type, img)
}
