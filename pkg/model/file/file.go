package file

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/model"
)

type File struct {
	gorm.Model
	Reference string
	System    string
	Type      string
	Path      string
	Name      string
}

func (f File) Add() (File, error) {
	model.Migrate(&f)
	if err := model.DB.Save(&f).Error; err != nil {
		return File{}, err
	}
	return f, nil
}

func Get(maps interface{}) ([]*File, error) {
	var files []*File
	err := model.DB.Where(maps).Find(&files).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return files, nil
}
