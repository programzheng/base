package model

import (
	"errors"

	_ "github.com/programzheng/base/config"

	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	globalDB *gorm.DB
)

func init() {
	var err error
	//?parseTime=true for the database table column type is TIMESTAMP
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&loc=Local&parseTime=true",
		viper.Get("DB_USERNAME").(string),
		viper.Get("DB_PASSWORD").(string),
		viper.Get("DB_HOST").(string),
		viper.Get("DB_PORT").(string),
		viper.Get("DB_DATABASE"))
	fmt.Printf("connect: %v database\n", dsn)
	globalDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		log.Println("DataBase error:", err)
	}
}

func GetDB() *gorm.DB {
	return globalDB
}

func HasTable(dst interface{}) bool {
	return GetDB().Migrator().HasTable(dst)
}

func CreateTable(dst ...interface{}) error {
	return GetDB().Migrator().CreateTable(dst...)
}

func SetupTableModel(models ...interface{}) error {
	//env is local
	if viper.Get("APP_ENV").(string) == "local" {
		err := GetDB().AutoMigrate(models...)
		if err != nil {
			log.Fatal(err)
		}
		return err
	}

	return errors.New("")
}
