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

func Get(post Post) (result interface{}) {
	result = model.Get(&post)
	return
}

func Add(post Post) (result interface{}) {
	result = model.Add(&post)
	return
}
