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

type Result struct {
	Value interface{}
	Error error
}

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

//Get first model to database
func Get(model interface{}) (result interface{}) {
	result = db.First(model)
	return
}

//Add model to database
// interface can't get origin variable only get variable at memory location
func Add(model interface{}) (result Result) {
	//create table for the struct
	result = Result{
		Value: nil,
		Error: nil,
	}
	db.AutoMigrate(model)
	if dbc := db.Create(model); dbc.Error != nil {
		//error
		result.Error = dbc.Error
	}
	return
}
