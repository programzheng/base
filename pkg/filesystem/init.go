package filesystem

import (
	"context"
	"io"
	"path/filepath"

	"github.com/programzheng/base/pkg/helper"

	"github.com/spf13/viper"
)

type FileSystem interface {
	Check()
	GetSystem() string
	GetPath() string
	Upload(context.Context, string, io.Reader) *StaticFile
	GetHostURL() string
}

type StaticFile struct {
	Reference   *string
	System      string
	Type        string
	Path        string
	Name        string
	ThirdPatyID string
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

func GetEmptyImageLink() string {
	return "//" + viper.Get("APP_URL").(string) + ":" + viper.Get("APP_PORT").(string) + "/files/image/empty"
}
