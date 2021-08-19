package admin

import (
	"base/pkg/model/admin"

	"github.com/jinzhu/copier"
)

type Admin struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	GroupID  int
	Status   int
	Profile  AdminProfile `json:"profile"`

	PageNum  int `form:"page_num"`  //頁數*筆數,從0(代表第一頁)開始
	PageSize int `form:"page_size"` //從PageNum之後取出的筆數
}

type AdminProfile struct {
	AdminID uint
	Name    string `json:"name"`
}

func (a *Admin) Add() (Admin, error) {
	modelAdmin := admin.Admin{
		Account:  a.Account,
		Password: a.Password,
		Profile: admin.AdminProfile{
			Name: a.Profile.Name,
		},
	}

	result, err := modelAdmin.Add()
	if err != nil {
		return Admin{}, err
	}

	admin := Admin{}
	copier.Copy(&admin, &result)

	return admin, nil
}

func (a *Admin) GetForLogin() (*admin.Admin, error) {
	modelAdmin := admin.Admin{
		Account:  a.Account,
		Password: a.Password,
	}
	model, err := modelAdmin.GetForLogin()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (a *Admin) Get() ([]*admin.Admin, error) {
	models, err := admin.Get(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (a *Admin) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
