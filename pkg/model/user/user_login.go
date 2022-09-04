package user

import (
	"github.com/programzheng/base/pkg/model"
	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	UserID   uint
	User     User
	Token    string
	ClientIP *string
}

func (ul *UserLogin) Add() (*UserLogin, error) {
	if err := model.GetDB().Create(&ul).Error; err != nil {
		return nil, err
	}

	return ul, nil
}

func (ul *UserLogin) First() (*UserLogin, error) {
	if err := model.GetDB().Joins("User").Where(&ul).First(&ul).Error; err != nil {
		return nil, err
	}

	return ul, nil
}

func (ul *UserLogin) Update(where map[string]interface{}, update map[string]interface{}) (*UserLogin, error) {
	if err := model.GetDB().Model(&ul).Where(where).Updates(update).Error; err != nil {
		return nil, err
	}

	return ul, nil
}

func (ul *UserLogin) BatchForceDelete(where map[string]interface{}) (int, error) {
	tx := model.GetDB().Where(where).Unscoped().Delete(&ul)
	if err := tx.Error; err != nil {
		return 0, err
	}

	return int(tx.RowsAffected), nil
}
