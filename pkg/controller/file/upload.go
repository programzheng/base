package file

import (
	"errors"
	"fmt"
	"path/filepath"

	"base/pkg/filesystem"
	"base/pkg/helper"
	"base/pkg/service/file"

	"github.com/gin-gonic/gin"
)

func init() {
	filesystem.Driver.Check()
}

func Upload(ctx *gin.Context) {
	//取得所有Mulitpart Form
	form, err := ctx.MultipartForm()
	if err != nil {
		helper.BadRequest(ctx, errors.New(fmt.Sprintf("get form error: %s", err.Error())))
		return
	}
	//取得所有File map[]
	uploadFileList := form.File
	//因為這樣取出來還會有一層map[]所以只能跑兩次
	fileList := file.Files{}
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
			fileType := helper.GetFileContentType(fileExtension)
			//上傳檔案
			if err := filesystem.Driver.Upload(ctx, uploadFile); err != nil {
				helper.BadRequest(ctx, errors.New(fmt.Sprintf("upload file err: %s", err.Error())))
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
				helper.BadRequest(ctx, errors.New(fmt.Sprintf("add file row error: %s", err.Error())))
				return
			}
			fileList = append(fileList, file)
		}
	}
	helper.UploadSuccess(ctx, fileList, "上傳成功")
}
