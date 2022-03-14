package file

import (
	"time"

	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/model"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	HashID      string  `gorm:"unique"`
	Reference   *string `json:"-"`
	System      string  `json:"-"`
	Type        string
	Path        string `json:"-"`
	Name        string `json:"-"`
	ThirdPatyID string `json:"-"`
}

type Files []*File

func init() {
	if !model.HasTable(&File{}) {
		model.CreateTable(&File{})
	}
}

func (f *File) AfterCreate(tx *gorm.DB) (err error) {
	// 設定給前端呼叫圖片的ID
	hashID := helper.ConvertToString(f.ID) + "_" + helper.ConvertToString(time.Now().Unix())
	hashID = helper.CreateMD5(hashID)
	tx.Model(f).Update("HashID", hashID)
	return
}

func (f File) Add() (File, error) {
	model.SetupTableModel(&f)
	if err := model.GetDB().Save(&f).Error; err != nil {
		return File{}, err
	}
	return f, nil
}

func Get(ids []interface{}, maps interface{}) (Files, error) {
	var files Files
	if ids == nil {
		err := model.GetDB().Where(maps).Find(&files).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return files, nil
	}
	err := model.GetDB().Where(ids).Where(maps).Find(&files).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return files, nil
}

func BatchUpdates(fn func() map[string]interface{}, updates interface{}) (Files, error) {
	maps := fn()
	var files Files
	err := model.GetDB().Model(&files).Where(maps).Updates(updates).Find(&files).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// err := model.GetDB().Model(&files).Where(ids).Where(maps).Updates(updates).Find(&files).Error

	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return nil, err
	// }

	return files, nil
}

func (f File) Update() (File, error) {
	if err := model.GetDB().Save(&f).Error; err != nil {
		return File{}, err
	}
	return f, nil
}
