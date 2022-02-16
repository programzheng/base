package filesystem

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/file"

	log "github.com/sirupsen/logrus"
)

type Local struct {
	System string
	Path   string
}

func (lSystem *Local) Check() {
	//檢查local路徑的資料夾有沒有存在
	if _, err := os.Stat(lSystem.Path); os.IsNotExist(err) {
		//make nested directories
		err = os.MkdirAll(lSystem.Path, 0700)
		if err != nil {
			log.Println("File system create directory error:", err)
			return
		}
	}
}

func (lSystem *Local) GetSystem() string {
	return lSystem.System
}

func (lSystem *Local) GetPath() string {
	return lSystem.Path
}

func (lSystem *Local) Upload(ctx *gin.Context, uploadFile *multipart.FileHeader) *file.File {
	//檔案位置
	filePath := lSystem.GetPath()
	//檔案名稱
	fileName := filepath.Base(uploadFile.Filename)
	//檔案副檔名
	fileExtension := filepath.Ext(fileName)
	//檔案mimeType
	fileType := helper.GetFileContentType(fileExtension)

	//利用gin的上傳檔案function
	err := ctx.SaveUploadedFile(uploadFile, filePath+"/"+fileName)
	if err != nil {
		log.Printf("filesystem local upload error:%v", err)
		return nil
	}

	fileService := file.File{
		System: lSystem.GetSystem(),
		Type:   fileType,
		Path:   filePath,
		Name:   fileName,
	}

	return &fileService
}

func (lSystem *Local) GetHostURL() string {
	return ""
}

func (lSystem *Local) GetLink(file file.File) string {
	return file.Path + "/" + file.Name
}
