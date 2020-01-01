package filesystem

import (
	"github.com/spf13/viper"
)

type FileSystem struct {
	System string
	Path   string
}

var Driver FileSystem

func init() {
	//driver value
	Driver.System = viper.Get("FILESYSTEM_DRIVER").(string)
	switch Driver.System {
	case "local":
		Driver.Path = viper.Get("FILESYSTEM_LOCAL_PATH").(string)
	}
}
