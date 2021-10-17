package filesystem

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type FileSystem interface {
	Check()
	GetSystem() string
	GetPath() string
	Upload(*gin.Context, *multipart.FileHeader) error
	GetHostURL() string
}

var (
	ex, _    = os.Executable()
	basepath = filepath.Dir(ex)
)

var Driver FileSystem

func init() {
	system := viper.Get("FILESYSTEM_DRIVER").(string)
	switch system {
	case "local":
		Driver = Local{
			System: viper.Get("FILESYSTEM_DRIVER").(string),
			Path:   filepath.Join(basepath, viper.Get("FILESYSTEM_LOCAL_PATH").(string)),
		}
	}
}
