package bot

import (
	"base/pkg/model"

	"github.com/jinzhu/gorm"
)

type LineBilling struct {
	gorm.Model
	BillingID uint `gorm:"unique; not null"`
	GroupID   string
	RoomID    string
	UserID    string `gorm:"not null"`
}

func init() {
	if !model.DB.HasTable(&LineBilling{}) {
		model.DB.CreateTable(&LineBilling{})
	}
}

func (lb LineBilling) Add() (LineBilling, error) {
	if err := model.DB.Create(&lb).Error; err != nil {
		return LineBilling{}, err
	}

	return lb, nil
}
