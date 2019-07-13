package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func environment() {
	switch os.Getenv("APP_ENV") {
	case "production":
		viper.SetConfigType("yaml")
		viper.SetConfigName("production")
		viper.AddConfigPath("./yaml")
	default:
		viper.SetConfigType("yaml")
		viper.SetConfigName("local")
		viper.AddConfigPath("./yaml")
	}
}

func init() {
	environment()
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
