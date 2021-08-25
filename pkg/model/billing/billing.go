package billing

import (
	"base/pkg/model"

	"github.com/jinzhu/gorm"
)

type Billing struct {
	gorm.Model
	Title  string `gorm:"comment:'標題'"`
	Amount int    `gorm:"comment:'總付款金額'"`
	Payer  string `gorm:"comment:'付款人'"`
	Note   string `gorm:"comment:'備註'"`
}

func init() {
	if !model.DB.HasTable(&Billing{}) {
		model.DB.CreateTable(&Billing{})
	}
}

func (b Billing) Add() (Billing, error) {
	if err := model.DB.Create(&b).Error; err != nil {
		return Billing{}, err
	}

	return b, nil
}
