package post

import (
	"github.com/ProgramZheng/base/model"
)

type Post struct {
	Id         int
	Title      string `json:title`
	Text       string `json:text`
	CreateTime int64
	UpdateTime int64
}

func Get(post Post) (result Post) {
	result = model.Get(&post)
	return
}

func Add(post Post) (status bool, err error) {
	status, err = model.Add(&post)
	return
}
