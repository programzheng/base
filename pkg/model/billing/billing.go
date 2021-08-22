package billing

import (
	"base/pkg/model"
	"time"
)

type Tabler interface {
	TableName() string
}

type Billing struct {
	model.Model
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"comment:'標題'"`
	Amount    int    `gorm:"comment:'總付款金額'"`
	Payer     string `gorm:"comment:'付款人'"`
	Note      string `gorm:"comment:'備註'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Billing) TableName() string {
	return "billings"
}

// func (b Billing) init() {
// 	if !model.DB.HasTable(&b) {
// 		model.DB.CreateTable(&b)
// 	}
// }

// func (b Billing) Add() (interface{}, error) {
// 	if err := model.DB.Create(&b).Error; err != nil {
// 		return Billing{}, err
// 	}

// 	return b, nil
// }
