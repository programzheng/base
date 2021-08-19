package post

import (
	"base/pkg/model"
	"base/pkg/model/file"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title         string
	Summary       string
	Detail        string `sql:"type:text"`
	FileReference string
	Files         []file.File `gorm:"unique;foreignkey:Reference;association_foreignkey:FileReference"`
}

func init() {
	if !model.DB.HasTable(&Post{}) {
		model.DB.CreateTable(&Post{})
	}
}

func (p Post) Add() (Post, error) {
	if err := model.DB.Create(&p).Error; err != nil {
		return Post{}, err
	}

	return p, nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Post, error) {
	var posts []*Post
	err := model.DB.Preload("Files").Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return posts, nil
}
