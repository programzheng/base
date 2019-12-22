package filesystem

import (
	"log"

	"github.com/programzheng/base/pkg/function"
	"github.com/spf13/viper"
)

var driver struct {
	Name string
	Path string
}

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
	driver.Name = viper.Get("FILESYSTEM_DRIVER").(string)
	driver.Path = fileSystemViper.Get(driver.Name + ".path").(string)
	function.GetStruct(driver)
}
