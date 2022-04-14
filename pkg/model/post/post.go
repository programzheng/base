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

func GetTotalNumber(maps interface{}) (int64, error) {
	var count int64

	err := model.GetDB().Model(&Post{}).Where(maps).Count(&count).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, nil
}

func Get(offset int, limit int, maps interface{}) ([]*Post, error) {
	var posts []*Post
	err := model.GetDB().Preload("Files").Where(maps).Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return posts, nil
}

func FindById(id uint) (*Post, error) {
	var post *Post
	err := model.GetDB().Preload("Files").First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func Update(maps interface{}, updates interface{}) (*Post, error) {
	var post Post
	err := model.GetDB().Model(&post).Where(maps).Updates(updates).Find(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func UpdateByModel(m Post, updates interface{}) (*Post, error) {
	var post Post
	err := model.GetDB().Model(&m).Updates(updates).Find(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func DelByID(id uint) error {
	var post Post
	err := model.GetDB().First(&post, id).Error
	if err != nil {
		return err
	}
	err = model.GetDB().Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}
