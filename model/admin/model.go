package admin

import (
	"github.com/ProgramZheng/base/model"
	"github.com/jinzhu/gorm"
)

//Login is vaildation struct
type Login struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Admin struct {
	gorm.Model
	Account  string `json:"account" gorm:"unique; not null"`
	Password string `json:"password" gorm:"unique; not null"`
	GroupID  int    `gorm:"index"`
	Status   int    `gorm:"defalut:0"`
	Profile  AdminProfile
}

type AdminProfile struct {
	gorm.Model
	AdminID uint
	Name    string `json:"name"`
}

func Add(admin Admin) (value interface{}, err error) {
	value, err = model.Add(&admin)
	return
}

func GetForID(admin Admin, id int) (result *Admin, err error) {
	value, err := model.GetForID(&admin, id)
	result = value.(*Admin)
	return
}

func Get(admin Admin, where interface{}) (result *Admin, err error) {
	value, err := model.Get(&admin, where)
	result = value.(*Admin)
	return
}

func SaveForID(admin Admin, id int, update interface{}) (result *Admin, err error) {
	value, err := model.SaveForID(&admin, id, update)
	result = value.(*Admin)
	return
}

func Save(admin Admin, ids interface{}, update map[string]interface{}) (result interface{}, err error) {
	result, err = model.Save(&admin, ids, update)
	return
}

func DelForID(admin Admin, id int) (result interface{}, err error) {
	result, err = model.DelForID(&admin, id)
	return
}

func Del(admin Admin, ids interface{}) (result interface{}, err error) {
	result, err = model.Del(&admin, ids)
	return
}
