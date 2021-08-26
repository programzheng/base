package bot

import (
	"base/pkg/model"
	"base/pkg/model/billing"

	"github.com/jinzhu/gorm"
)

type LineBilling struct {
	gorm.Model
	BillingID uint `gorm:"unique; not null"`
	GroupID   string
	RoomID    string
	UserID    string `gorm:"not null"`
	Billing   billing.Billing
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

func (lb LineBilling) Get(maps map[string]interface{}, not map[string]interface{}) ([]LineBilling, error) {
	var lbs []LineBilling
	err := model.DB.Preload("Billing").Where(maps).Not(not).Find(&lbs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return lbs, nil
}
