package bot

import (
	"base/pkg/model/bot"
	"base/pkg/service"
	"base/pkg/service/billing"

	"github.com/jinzhu/copier"
)

var (
	module string = "line_billing"
)

type LineBilling struct {
	BillingID uint
	GroupID   string
	RoomID    string
	UserID    string

	Billing billing.Billing

	Page service.Page
}

type LineBillings []LineBilling

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

func (lb *LineBilling) Get() (LineBillings, error) {
	results, err := bot.LineBilling{}.Get(lb.getMaps())
	if err != nil {
		return nil, err
	}
	lineBillings := LineBillings{}
	copier.Copy(&lineBillings, &results)
	return lineBillings, nil
}

func (lb *LineBilling) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
