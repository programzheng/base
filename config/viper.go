package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../")
)

type Instance struct {
	Package *viper.Viper
}

func loadEnv(v *viper.Viper) {
	envFilePath := filepath.Join(RootPath, ".env")
	//check .env is exist
	_, err := os.Stat(envFilePath)
	if !os.IsNotExist(err) {
		v.SetConfigFile(envFilePath)
		err = v.ReadInConfig() // Find and read the config file
		if err != nil {        // Handle errors reading the config file
			log.Fatalf("Fatal error config:%v", err)
		}
	}
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()
}

func NewViper() *viper.Viper {
	v := viper.New()
	loadEnv(v)

	return v
}

func (i *Instance) GetString(name string) string {
	return i.Package.GetString(name)
}
