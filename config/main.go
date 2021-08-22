package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

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
		path := "." + viper.Get("LOG_PATH").(string) + "/"
		//check log path directory exist
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0644)
			if err != nil {
				log.Fatal("cerate log directory error:", err)
			}
		}
		fileName := time.Now().Format("20060102") + ".log"
		file, err := os.OpenFile(path+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Info("Failed to log to file, using default stderr", err)
		}
		log.SetOutput(file)
		log.Info("log service running")
	}
}

func setViper() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Error("Fatal error config: ", err)
	}
}
