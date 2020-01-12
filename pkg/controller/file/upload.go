package file

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/service/file"
)

func init() {
	filesystem.Driver.Check()
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
	fileList := []file.ResponseFile{}
	//TODO: 調整迴圈
	for _, uploadFiles := range uploadFileList {
		for _, uploadFile := range uploadFiles {
			//檔案位置
			filePath := filesystem.Driver.GetPath()
			//檔案名稱
			fileName := filepath.Base(uploadFile.Filename)
			//檔案副檔名
			fileExtension := filepath.Ext(fileName)
			//檔案mimeType
			fileType := function.GetFileContentType(fileExtension)
			//上傳檔案
			if err := filesystem.Driver.Upload(ctx, uploadFile); err != nil {
				function.BadRequest(ctx, errors.New(fmt.Sprintf("upload file err: %s", err.Error())))
				return
			}

			fileService := file.File{
				System: filesystem.Driver.GetSystem(),
				Type:   fileType,
				Path:   filePath,
				Name:   fileName,
			}
			file, err := fileService.Add()
			if err != nil {
				function.BadRequest(ctx, errors.New(fmt.Sprintf("add file row error: %s", err.Error())))
				return
			}
			fileList = append(fileList, file)
		}
	}
	function.UploadSuccess(ctx, fileList, "上傳成功")
}
