package file

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	hashID := ctx.Param("hash_id")
	fmt.Println(hashID)
}
