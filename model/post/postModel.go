package post

import (
	"github.com/ProgramZheng/base/model"
)

type Post struct {
	ID         int
	Title      string `json:title`
	Text       string `json:text`
	CreateTime int64
	UpdateTime int64
}

func Add(post Post) (result interface{}) {
	result = model.Add(&post)
	return
}

func GetForID(post Post, id int) (result interface{}) {
	result = model.GetForID(&post, id)
	return
}

func Get(post Post) (result interface{}) {
	result = model.Get(&post)
	return
}

func Save(post Post, update interface{}) (result interface{}) {
	result = model.Save(&post, update)
	return
}
