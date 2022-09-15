package filesystem

import (
	"context"
	"io"
	"path/filepath"

	"github.com/programzheng/base/config"
	"github.com/programzheng/base/pkg/helper"
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
		system = config.Cfg.GetString("FILESYSTEM_DRIVER")
	}
	var Driver FileSystem
	switch system {
	case "local":
		Driver = &Local{
			System: system,
			Path:   filepath.Join(helper.RootPath, config.Cfg.GetString("FILESYSTEM_LOCAL_PATH")),
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
	return "//" + config.Cfg.GetString("APP_URL") + "/files/image/empty"
}
