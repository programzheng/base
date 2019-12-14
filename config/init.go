package config

import (
	"log"

	//use this read .env
	_ "github.com/joho/godotenv/autoload"
	//use this get .env setting
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Using config file:", err)
	}
}
