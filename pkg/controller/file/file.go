package file

import (
	"errors"
	"log"
	"net/http"

	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	hashID := ctx.Param("hash_id")
	where := make(map[string]interface{}, 1)
	where["hash_id"] = hashID
	files, err := file.Get(where)

	if err != nil {
		log.Fatalf("files get error: %v", err)
	}
	if len(files) == 0 {
		resource.BadRequest(ctx, errors.New("files get not found"))
		return
	}
	file := files[len(files)-1]
	img, err := file.GetBytes()
	if err != nil {
		log.Fatal(err)
	}
	ctx.Data(http.StatusOK, file.Type, img)
}
