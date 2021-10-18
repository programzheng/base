package file

import (
	"base/pkg/service/file"
	"fmt"
	"log"

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
	fmt.Printf("%v\n", files)
}
