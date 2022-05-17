package admin

import (
	"github.com/programzheng/base/pkg/model"

	"gorm.io/gorm"
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

func (a Admin) Add() (Admin, error) {
	if err := model.GetDB().Create(&a).Error; err != nil {
		return Admin{}, err
	}

	return a, nil
}

func (a Admin) GetForLogin() (*Admin, error) {
	if err := model.GetDB().Where(&a).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Admin, error) {
	var models []*Admin
	err := model.GetDB().Preload("Profile").Where(maps).Offset(pageNum).Limit(pageSize).Find(&models).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return models, nil
}
