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

func (lb *LineBilling) Get(where map[string]interface{}, not map[string]interface{}) ([]LineBilling, error) {
	results, err := bot.LineBilling{}.Get(service.GetDefaultWhere(where), not)
	if err != nil {
		return nil, err
	}
	var lineBillings []LineBilling
	copier.Copy(&lineBillings, &results)
	return lineBillings, nil
}
