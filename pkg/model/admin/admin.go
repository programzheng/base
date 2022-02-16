package admin

import (
	"github.com/programzheng/base/pkg/model"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Account  string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null" json:"-"`
	Status   int    `gorm:"defalut:0"`
	Profile  AdminProfile
}

type AdminProfile struct {
	gorm.Model
	AdminID uint
	Name    string
}

func init() {
	if !model.DB.HasTable(&Admin{}) {
		model.DB.CreateTable(&Admin{})
	}
	if !model.DB.HasTable(&AdminProfile{}) {
		model.DB.CreateTable(&AdminProfile{})
	}
}

func (a Admin) Add() (Admin, error) {
	if err := model.DB.Create(&a).Error; err != nil {
		return Admin{}, err
	}

	return a, nil
}

func (a Admin) GetForLogin() (*Admin, error) {
	if err := model.DB.Where(&a).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Admin, error) {
	var models []*Admin
	err := model.DB.Preload("Profile").Where(maps).Offset(pageNum).Limit(pageSize).Find(&models).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return models, nil
}
