package filesystem

import (
	"log"

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
	switch Driver.Name {
	case "local":
		Driver.Path = viper.Get("FILESYSTEM_LOCAL_PATH").(string)
	}
}

// func driverSetValue() (Driver FileSystem) {

// }
