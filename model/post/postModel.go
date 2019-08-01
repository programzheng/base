package post

import (
	"fmt"

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
	fmt.Println(result)
	return
}

func Add(post Post) (status bool, err error) {
	status, err = model.Add(&post)
	return
}
