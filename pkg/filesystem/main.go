package filesystem

import (
	"log"

	"github.com/programzheng/base/pkg/function"
	"github.com/spf13/viper"
)

type FileSystem struct {
	Name string
	Path string
}

var Driver FileSystem

func init() {
	//get file system config
	fileSystemViper := viper.New()
	fileSystemViper.SetConfigType("yaml")
	fileSystemViper.AddConfigPath("config/yaml")
	fileSystemViper.SetConfigName("filesystem")
	err := fileSystemViper.ReadInConfig()
	if err != nil {
		log.Println("file system config error:", err)
	}
	//driver value
	Driver.Name = viper.Get("FILESYSTEM_DRIVER").(string)
	switch Driver.Name {
	case "local":
		Driver.Path = fileSystemViper.Get(Driver.Name + ".path").(string)
	}
	function.GetStruct(Driver)
}
