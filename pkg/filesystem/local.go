package filesystem

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/programzheng/base/pkg/helper"

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

func (lSystem *Local) Upload(ctx context.Context, originFileName string, uploadFile io.Reader) *StaticFile {
	//檔案位置
	filePath := lSystem.GetPath()
	//檔案名稱
	fileName := helper.CreateUuid()
	//檔案副檔名
	fileExtension := filepath.Ext(originFileName)
	//完整檔案名稱
	fileFullName := fileName + fileExtension
	//檔案mimeType
	fileType := helper.GetFileContentType(fileExtension)

	//利用gin的上傳檔案function
	out, err := os.Create(filePath + "/" + fileFullName)
	if err != nil {
		log.Printf("filesystem local upload error:%v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, uploadFile)
	if err != nil {
		log.Printf("filesystem local upload error:%v", err)
	}

	staticFile := StaticFile{
		System: lSystem.GetSystem(),
		Type:   fileType,
		Path:   filePath,
		Name:   fileFullName,
	}

	return &staticFile
}

func (lSystem *Local) GetHostURL() string {
	return ""
}
