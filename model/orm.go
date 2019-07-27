package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	setting := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_DATABASE"))

	db, err = gorm.Open("mysql", setting)

	// defer db.Close()
	if err != nil {
		log.Println("DataBase error:", err)
	}
	// err = db.Ping()
	// if err != nil {
	// 	log.Println("DataBase Ping error:", err)
	// }
}

func Add(tableName string, model interface{}) (result bool, err error) {
	fmt.Println(model)
	// db.AutoMigrate(&model{})
	db.Table(tableName).Create(&model)
	// result = db.NewRecord(model)
	return result, err
}
