package filesystem

import (
	"mime/multipart"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Local struct {
	System string
	Path   string
}

func (l Local) Check() {
	//檢查local路徑的資料夾有沒有存在
	if _, err := os.Stat(l.Path); os.IsNotExist(err) {
		//建立資料夾,權限設為0777(-rwxrwxrwx)
		os.MkdirAll(l.Path, os.ModePerm)
		if err != nil {
			log.Println("File system create directory error:", err)
			return
		}
	}
}

func (l Local) GetSystem() string {
	return l.System
}

func (l Local) GetPath() string {
	return l.Path
}

func (l Local) Upload(ctx *gin.Context, uploadFile *multipart.FileHeader) error {
	//檔案位置
	filePath := l.Path
	//檔案名稱
	fileName := filepath.Base(uploadFile.Filename)
	//利用gin的上傳檔案function
	err := ctx.SaveUploadedFile(uploadFile, filePath+"/"+fileName)
	return err
}

func (l Local) GetHostURL() string {
	return viper.Get("APP_URL").(string)
}
