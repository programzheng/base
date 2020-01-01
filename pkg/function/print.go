package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetPostJSON(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
}

func GetJSON(value interface{}) {
	result, _ := json.Marshal(value)
	fmt.Println(string(result))
}
