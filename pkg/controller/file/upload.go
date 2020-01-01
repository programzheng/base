package file

import (
	"errors"
	// "io"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "encoding/binary"

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
	//取得所有Mulitpart Form
	form, err := ctx.MultipartForm()
	if err != nil {
		function.BadRequest(ctx, errors.New(fmt.Sprintf("get form error: %s", err.Error())))
		return
	}
	//取得所有File map[]
	uploadFileList := form.File
	//因為這樣取出來還會有一層map[]所以只能跑兩次
	IDList := []uint{}
	//TODO: 調整迴圈
	for _, uploadFiles := range uploadFileList {
		for _, uploadFile := range uploadFiles {
			//檔案位置
			filePath := filesystem.Driver.Path
			//檔案名稱
			fileName := filepath.Base(uploadFile.Filename)
			//檔案副檔名
			fileExtension := filepath.Ext(fileName)
			//檔案mimeType
			fileType := function.GetFileContentType(fileExtension)
			//利用gin的上傳檔案function
			if err := ctx.SaveUploadedFile(uploadFile, filePath+"/"+fileName); err != nil {
				function.BadRequest(ctx, errors.New(fmt.Sprintf("upload file err: %s", err.Error())))
				return
			}

			fileService := file.File{
				System: filesystem.Driver.System,
				Type:   fileType,
				Path:   filePath,
				Name:   fileName,
			}
			ID, err := fileService.Add()
			if err != nil {
				function.BadRequest(ctx, errors.New(fmt.Sprintf("add file row error: %s", err.Error())))
				return
			}
			IDList = append(IDList, ID)
		}
	}
	function.UploadSuccess(ctx, IDList, "上傳成功")
}
