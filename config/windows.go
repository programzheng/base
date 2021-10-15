// +build windows

package config

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"

	//use this get .env setting
	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	setViper()
	setTimeZone()
	setLog()
}

func setViper() {
	envFilePath := filepath.Join(basepath, "../.env")
	//check .env is exist
	_, err := os.Stat(envFilePath)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFilePath)
		err = viper.ReadInConfig() // Find and read the config file
		if err != nil {            // Handle errors reading the config file
			log.Error("Fatal error config: ", err)
		}
	}
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
}

func setTimeZone() {
	TZ := viper.Get("TZ").(string)
	local, err := time.LoadLocation(TZ)
	if err != nil {
		log.Error("set timezone error:", err)
	}
	time.Local = local
}

func setLog() {
	system := viper.Get("LOG_SYSTEM").(string)
	switch system {
	case "file":
		path := filepath.Join(basepath, "../"+viper.Get("LOG_PATH").(string))
		//check log path directory exist
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			//make nested directories
			err = os.MkdirAll(path, 0700)
			if err != nil {
				log.Fatal("create log directory error:", err)
			}
		}
		fileName := time.Now().Format("20060102") + ".log"
		filePath := filepath.Join(path, fileName)
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
		if err != nil {
			log.Info("Failed to log to file, using default stderr", err)
		}
		log.SetOutput(file)
		logLevel := viper.Get("LOG_LEVEL").(string)
		logLevelNumber, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Info("log ParseLevel error:", err)
		}
		// Only log the warning severity or above.
		log.SetLevel(logLevelNumber)
		log.Info("log service running")
	}
}
