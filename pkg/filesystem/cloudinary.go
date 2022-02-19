package filesystem

import (
	"context"
	"io"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/programzheng/base/pkg/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cloudinary struct {
	System string
	Path   string
}

func getCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		viper.Get("FILESYSTEM_CLOUDINARY_CLOUD_NAME").(string),
		viper.Get("FILESYSTEM_CLOUDINARY_API_KEY").(string),
		viper.Get("FILESYSTEM_CLOUDINARY_SECRET").(string),
	)
	if err != nil {
		log.Printf("File system cloudinary getCloudinary error:%v", err)
	}
	return cld
}

func (cldSystem *Cloudinary) Check() {

}

func (cldSystem *Cloudinary) GetSystem() string {
	return cldSystem.System
}

func (cldSystem *Cloudinary) GetPath() string {
	return cldSystem.Path
}

func (cldSystem *Cloudinary) Upload(ctx context.Context, originFileName string, uploadFile io.Reader) *StaticFile {
	//檔案名稱
	fileName := helper.CreateUuid()
	//檔案副檔名
	fileExtension := filepath.Ext(originFileName)
	//完整檔案名稱
	fileFullName := fileName + fileExtension
	//檔案mimeType
	fileType := helper.GetFileContentType(fileExtension)

	cld := getCloudinary()
	resp, err := cld.Upload.Upload(ctx, uploadFile, uploader.UploadParams{})
	if err != nil {
		log.Printf("File system cloudinary upload error:%v", err)
	}

	//replace https to ""
	path := resp.SecureURL
	if strings.HasPrefix(path, "https:") {
		path = strings.Replace(path, "https:", "", 1)
	}

	staticFile := StaticFile{
		System:      cldSystem.GetSystem(),
		Type:        fileType,
		Path:        path,
		Name:        fileFullName,
		ThirdPatyID: resp.PublicID,
	}

	return &staticFile
}

func (cldSystem *Cloudinary) GetHostURL() string {
	return ""
}
