package post

import (
	"github.com/ProgramZheng/base/model"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:title`
	Text  string `json:text`
}

func Add(post Post) (result interface{}) {
	result = model.Add(&post)
	return
}

func GetForID(post Post, id int) (result interface{}) {
	result = model.GetForID(&post, id)
	return
}

func Get(post Post, where interface{}) (result interface{}) {
	result = model.Get(&post, where)
	return
}

func SaveForID(post Post, id int, update interface{}) (result interface{}) {
	result = model.Save(&post, id, update)
	return
}

func DelForID(post Post, id int) (result interface{}) {
	result = model.Del(&post, id)
	return
}
