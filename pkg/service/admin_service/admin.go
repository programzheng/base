package admin_service

import (
	"github.com/ProgramZheng/base/pkg/model"
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

func (a *Admin) Add() error {
	admin := model.Admin{
		Account:  a.Account,
		Password: a.Password,
		Profile: model.AdminProfile{
			Name: a.Profile.Name,
		},
	}
	if err := model.AddAdmin(admin); err != nil {
		return err
	}
	return nil
}

func (a *Admin) Get() (*model.Admin, error) {
	where := model.Admin{
		Account:  a.Account,
		Password: a.Password,
	}
	admin, err := model.GetAdmin(where)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (a *Admin) GetAdmins() ([]*model.Admin, error) {
	admins, err := model.GetAdmins(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (a *Admin) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
