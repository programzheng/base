package filesystem

import (
	"fmt"

	"github.com/spf13/viper"
)

var fileSystemDriver string

func init() {
	fileSystemDriver = viper.Get("FILESYSTEM_DRIVER").(string)
	fmt.Println(fileSystemDriver)
}
