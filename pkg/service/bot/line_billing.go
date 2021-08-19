package bot

import (
	"base/pkg/model/bot"
	"base/pkg/service"

	"github.com/jinzhu/copier"
)

type LineBilling struct {
	BillingID uint
	GroupID   string
	RoomID    string
	UserID    string

	service.Page
}

func (lb *LineBilling) Add() (LineBilling, error) {
	model := bot.LineBilling{}
	copier.Copy(&model, &lb)
	result, err := model.Add()
	if err != nil {
		return LineBilling{}, err
	}
	lineBilling := LineBilling{}
	copier.Copy(&lineBilling, &result)

	return lineBilling, nil
}
