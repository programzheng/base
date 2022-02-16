package filesystem

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cloudinary struct {
	System string
	Path   string
}

func getCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		viper.Get("FILESYSTEM_LOCAL_CLOUDINARY_CLOUD_NAME").(string),
		viper.Get("FILESYSTEM_LOCAL_CLOUDINARY_API_KEY").(string),
		viper.Get("FILESYSTEM_LOCAL_CLOUDINARY_SECRET").(string),
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

func (cldSystem *Cloudinary) Upload(ctx *gin.Context, uploadFile *multipart.FileHeader) *file.File {
	openFile, err := uploadFile.Open()
	if err != nil {
		log.Printf("File system cloudinary upload open() error:%v", err)
	}

	//檔案名稱
	fileName := filepath.Base(uploadFile.Filename)
	//檔案副檔名
	fileExtension := filepath.Ext(fileName)
	//檔案mimeType
	fileType := helper.GetFileContentType(fileExtension)

	cld := getCloudinary()
	resp, err := cld.Upload.Upload(ctx, openFile, uploader.UploadParams{})
	if err != nil {
		log.Printf("File system cloudinary upload error:%v", err)
	}

	//replace https to ""
	path := resp.SecureURL
	if strings.HasPrefix(path, "https:") {
		path = strings.Replace(path, "https:", "", 1)
	}

	fileService := file.File{
		Reference: resp.PublicID,
		System:    cldSystem.GetSystem(),
		Type:      fileType,
		Path:      path,
		Name:      fileName,
	}

	return &fileService
}

func (cldSystem *Cloudinary) GetHostURL() string {
	return ""
}

func (cldSystem *Cloudinary) GetLink(file file.File) string {
	return file.Path
}
