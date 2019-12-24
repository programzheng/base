package upload

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/function"
)

func init() {
	switch filesystem.Driver.Name {
	case "local":
		if _, err := os.Stat(filesystem.Driver.Path); os.IsNotExist(err) {
			os.MkdirAll(filesystem.Driver.Path, os.ModePerm)
			if err != nil {
				log.Println("File system create directory error:", err)
				return
			}
		}
	}
}

func File(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("upload")
	if err != nil {
		function.Fail(ctx, err)
		return
	}
	function.GetStruct(filesystem.Driver)
	//檔案位置
	filePath := filesystem.Driver.Path
	//檔案名稱
	fileName := header.Filename
	fmt.Println(filePath + "/" + fileName)
	//建立空檔案
	out, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		function.Fail(ctx, err)
	}
	defer out.Close()

	//複製檔案內容
	_, err = io.Copy(out, file)
	if err != nil {
		function.Fail(ctx, err)
	}

	function.UploadSuccess(ctx, nil, "上傳成功")
}
