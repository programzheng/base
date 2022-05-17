package model

import (
	"errors"

	"github.com/programzheng/base/config"

	"fmt"

	log "github.com/sirupsen/logrus"

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
		config.Cfg.GetString("DB_USERNAME"),
		config.Cfg.GetString("DB_PASSWORD"),
		config.Cfg.GetString("DB_HOST"),
		config.Cfg.GetString("DB_PORT"),
		config.Cfg.GetString("DB_DATABASE"))
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
	if config.Cfg.GetString("APP_ENV") == "local" {
		err := GetDB().AutoMigrate(models...)
		if err != nil {
			log.Fatal(err)
		}
		return err
	}

	return errors.New("")
}
