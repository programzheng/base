package post

import (
	"github.com/programzheng/base/pkg/model"
	"github.com/programzheng/base/pkg/model/file"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string
	Summary       string
	Detail        string      `gorm:"type:longtext"`
	FileReference *string     `gorm:"unique"`
	Files         []file.File `gorm:"foreignKey:Reference;references:FileReference;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func init() {
	if !model.HasTable(&Post{}) {
		model.CreateTable(&Post{})
	}
}

func (p Post) Add() (Post, error) {
	if err := model.GetDB().Create(&p).Error; err != nil {
		return Post{}, err
	}

	return p, nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Post, error) {
	var posts []*Post
	err := model.GetDB().Preload("Files").Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return posts, nil
}
