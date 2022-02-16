package filesystem

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/spf13/viper"
)

type FileSystem interface {
	Check()
	GetLink(file file.File) string
	GetSystem() string
	GetPath() string
	Upload(*gin.Context, *multipart.FileHeader) *file.File
	GetHostURL() string
}

func Create(system string) FileSystem {
	if system == "" {
		system = viper.Get("FILESYSTEM_DRIVER").(string)
	}
	var Driver FileSystem
	switch system {
	case "local":
		Driver = &Local{
			System: system,
			Path:   filepath.Join(helper.RootPath, viper.Get("FILESYSTEM_LOCAL_PATH").(string)),
		}
	case "cloudinary":
		Driver = &Cloudinary{
			System: system,
			Path:   "",
		}
	}

	return Driver
}
