package model

import (
	"github.com/ProgramZheng/base/model"
)

const tableName = "post"

type Post struct {
	Id         int    `gorm:"AUTO_INCREMENT"`
	Title      string `json:title`
	Text       string `json:text`
	CreateTime int64
	UpdateTime int64
}

func Add(post Post) {
	model.Add(tableName, post)
}
