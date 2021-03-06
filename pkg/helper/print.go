package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetType(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}

func GetPostJSON(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
}

func GetJSON(value interface{}) {
	result, _ := json.Marshal(value)
	fmt.Println(string(result))
}

func GetGinRequest(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	fmt.Println(reqBody)
}
