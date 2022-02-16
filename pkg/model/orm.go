package model

import (
	_ "github.com/programzheng/base/config"
	"github.com/programzheng/base/pkg/helper"

	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	//?parseTime=true for the database table column type is TIMESTAMP
	setting := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?loc=Local&parseTime=true",
		viper.Get("DB_USERNAME").(string),
		viper.Get("DB_PASSWORD").(string),
		viper.Get("DB_HOST").(string),
		viper.Get("DB_PORT").(string),
		viper.Get("DB_DATABASE"))
	fmt.Printf("connect: %v database\n", setting)
	DB, err = gorm.Open(viper.Get("DB_CONNECTION").(string), setting)
	DB.LogMode(helper.ConvertToBool(viper.Get("APP_debug").(string)))

	if err != nil {
		log.Println("DataBase error:", err)
	}
}

func Migrate(models ...interface{}) {
	DB.AutoMigrate(models...)
}
