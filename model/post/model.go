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

func Add(post Post) (value interface{}, err error) {
	value, err = model.Add(&post)
	return
}

func GetForID(post Post, id int) (result *Post, err error) {
	value, err := model.GetForID(&post, id)
	result = value.(*Post)
	return
}

func Get(post Post, where interface{}) (result *Post, err error) {
	value, err := model.Get(&post, where)
	result = value.(*Post)
	return
}

func SaveForID(post Post, id int, update interface{}) (result *Post, err error) {
	value, err := model.SaveForID(&post, id, update)
	result = value.(*Post)
	return
}

func Save(post Post, ids interface{}, update map[string]interface{}) (result interface{}, err error) {
	result, err = model.Save(&post, ids, update)
	return
}

func DelForID(post Post, id int) (result interface{}, err error) {
	result, err = model.DelForID(&post, id)
	return
}

func Del(post Post, ids interface{}) (result interface{}, err error) {
	result, err = model.Del(&post, ids)
	return
}
