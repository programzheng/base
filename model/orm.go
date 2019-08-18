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
	//?parseTime=true for the database table column type is TIMESTAMP
	setting := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?loc=Local&parseTime=true", viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_DATABASE"))

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

//Add is add model to database
// interface can't get origin variable only get variable at memory location
func Add(model interface{}) (value interface{}, err error) {
	//create table for the struct
	db.AutoMigrate(model)
	if dbc := db.Create(model); dbc.Error != nil {
		//error
		value = nil
		err = dbc.Error
	}
	return
}

//GetForID is get first model For ID from database
func GetForID(model interface{}, id int) (value interface{}, err error) {
	result := db.First(model, id)
	value = result.Value
	err = result.Error
	return
}

//Get is get first model to database
func Get(model interface{}, where interface{}) (value interface{}, err error) {
	result := db.Where(where).Find(model)
	value = result.Value
	err = result.Error
	return
}

//SaveForID is get first model then update data to database
func SaveForID(model interface{}, id int, update interface{}) (value interface{}, err error) {
	result := db.First(model, id).Updates(update)
	value = result.Value
	err = result.Error
	return
}

//Save is get ID's list For table name then update data to database
func Save(model interface{}, ids interface{}, update interface{}) (value interface{}, err error) {
	result := db.Table(db.NewScope(model).TableName()).Where("id IN (?)", ids).Updates(update)
	value = result.Value
	err = result.Error
	return
}

//DelForID is del find model to database
func DelForID(model interface{}, id int) (value interface{}, err error) {
	result := db.Delete(model, id)
	value = result.Value
	err = result.Error
	return
}

//Del is get ID's list For table name then delete data to database
func Del(model interface{}, ids interface{}) (value interface{}, err error) {
	result := db.Table(db.NewScope(model).TableName()).Where("id IN (?)", ids).Delete(model)
	value = result.Value
	err = result.Error
	return
}
