package admin

import (
	"github.com/ProgramZheng/base/model"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Account  string `json:account`
	Password string `json:password`
	GroupID  int
}

func Add(admin Admin) (result interface{}) {
	result = model.Add(&admin)
	return
}

func GetForID(admin Admin, id int) (result interface{}) {
	result = model.GetForID(&admin, id)
	return
}

func Get(admin Admin, where interface{}) (result interface{}) {
	result = model.Get(&admin, where)
	return
}

func SaveForID(admin Admin, id int, update interface{}) (result interface{}) {
	result = model.SaveForID(&admin, id, update)
	return
}

func Save(admin Admin, ids interface{}, update map[string]interface{}) (result interface{}) {
	result = model.Save(&admin, ids, update)
	return
}

func DelForID(admin Admin, id int) (result interface{}) {
	result = model.Del(&admin, id)
	return
}

func Del(admin Admin, ids interface{}) (result interface{}) {
	result = model.Del(&admin, ids)
	return
}
