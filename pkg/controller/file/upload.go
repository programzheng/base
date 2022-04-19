package file

import (
	"errors"
	"fmt"

	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/gin-gonic/gin"
)

func init() {
	filesystem.Create("").Check()
}

func Upload(ctx *gin.Context) {
	//取得所有Mulitpart Form
	form, err := ctx.MultipartForm()
	if err != nil {
		resource.BadRequest(ctx, errors.New(fmt.Sprintf("get form error: %s", err.Error())))
		return
	}
	//取得所有File map[]
	uploadFileList := form.File
	//因為這樣取出來還會有一層map[]所以只能跑兩次
	fileList := file.Files{}
	//TODO: 調整迴圈
	for _, uploadFiles := range uploadFileList {
		for _, uploadFile := range uploadFiles {
			originFileName := uploadFile.Filename
			uf, err := uploadFile.Open()
			if err != nil {
				resource.BadRequest(ctx, errors.New(fmt.Sprintf("upload error: %s", err.Error())))
				return
			}
			//上傳檔案
			staticFile := filesystem.Create("").Upload(ctx, originFileName, uf)
			if staticFile == nil {
				resource.BadRequest(ctx, errors.New("upload file error"))
				return
			}
			fileService := file.File{
				Reference:   staticFile.Reference,
				System:      staticFile.System,
				Type:        staticFile.Type,
				Path:        staticFile.Path,
				Name:        staticFile.Name,
				ThirdPatyID: staticFile.ThirdPatyID,
			}

			file, err := fileService.Add()
			if err != nil {
				resource.BadRequest(ctx, fmt.Errorf("add file row error: %v", err))
				return
			}
			fileList = append(fileList, file)
		}
	}
	resource.UploadSuccess(ctx, fileList, "上傳成功")
}
