package upload

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func File(ctx *gin.Context) {
	fmt.Println(ctx)
}
