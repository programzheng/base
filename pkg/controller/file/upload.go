package file

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/service/file"
)

func init() {
	switch filesystem.Driver.System {
	case "local":
		//檢查local路徑的資料夾有沒有存在
		if _, err := os.Stat(filesystem.Driver.Path); os.IsNotExist(err) {
			//建立資料夾,權限設為0777(-rwxrwxrwx)
			os.MkdirAll(filesystem.Driver.Path, os.ModePerm)
			if err != nil {
				log.Println("File system create directory error:", err)
				return
			}
		}
	}
}

func Upload(ctx *gin.Context) {
	uploadFile, header, err := ctx.Request.FormFile("upload")
	if err != nil {
		function.Fail(ctx, err)
		return
	}
	//檔案位置
	filePath := filesystem.Driver.Path
	//檔案名稱
	fileName := header.Filename
	//檔案副檔名
	fileExtension := filepath.Ext(fileName)
	//建立空檔案
	out, err := os.Create(filePath + "/" + fileName)
	if err != nil {
		function.Fail(ctx, err)
		return
	}
	defer out.Close()

	//複製檔案內容
	_, err = io.Copy(out, uploadFile)
	if err != nil {
		function.Fail(ctx, err)
		return
	}

	fileService := file.File{
		System: filesystem.Driver.System,
		Type:   function.GetFileContentType(fileExtension),
		Path:   filePath,
		Name:   fileName,
	}

	if err := fileService.Add(); err != nil {
		function.Fail(ctx, err)
		return
	}

	function.UploadSuccess(ctx, nil, "上傳成功")
}
