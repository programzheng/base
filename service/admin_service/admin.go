package admin_service

import (
	"github.com/ProgramZheng/base/model"
)

type Admin struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	GroupID  int
	Status   int
	Profile  AdminProfile `json:"profile"`

	PageNum  int
	PageSize int
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
