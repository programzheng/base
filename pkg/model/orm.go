package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	//?parseTime=true for the database table column type is TIMESTAMP
	setting := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?loc=Local&parseTime=true",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_DATABASE"))
	db, err = gorm.Open(viper.GetString("DB_CONNECTION"), setting)

	if err != nil {
		log.Println("DataBase error:", err)
	}
}

func Migrate(models ...interface{}) {
	db.AutoMigrate(models...)
}
