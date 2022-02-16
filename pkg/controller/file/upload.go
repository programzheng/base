package file

import (
	"errors"
	"fmt"

	"base/pkg/filesystem"
	"base/pkg/helper"
	"base/pkg/service/file"

	"github.com/gin-gonic/gin"
)

func init() {
	filesystem.Create("").Check()
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
			//上傳檔案
			fileService := filesystem.Create("").Upload(ctx, uploadFile)
			if fileService == nil {
				helper.BadRequest(ctx, errors.New("upload file error"))
				return
			}

			file, err := fileService.Add()
			if err != nil {
				helper.BadRequest(ctx, fmt.Errorf("add file row error: %v", err))
				return
			}
			fileList = append(fileList, file)
		}
	}
	helper.UploadSuccess(ctx, fileList, "上傳成功")
}
