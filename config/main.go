package config

import (

	//use this read .env

	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"

	//use this get .env setting
	"github.com/spf13/viper"
)

func init() {
	setViper()
	setTimeZone()
	setLog()
}

func setTimeZone() {
	TZ := viper.Get("TZ").(string)
	os.Setenv("TZ", TZ)
}

func setLog() {
	system := viper.Get("LOG_SYSTEM").(string)
	switch system {
	case "file":
		path := viper.Get("LOG_PATH").(string) + "/"
		fileName := time.Now().Format("20060102") + ".log"
		file, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			logrus.Info("Failed to log to file, using default stderr")
		}
		logrus.Info("log service running")
	}
}

func setViper() {
	viper.AutomaticEnv()
}
