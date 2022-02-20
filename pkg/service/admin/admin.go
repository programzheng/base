package admin

import (
	"github.com/programzheng/base/pkg/model/admin"
	"github.com/programzheng/base/pkg/service"

	"github.com/jinzhu/copier"
)

type Admin struct {
	Account  string       `json:"account"`
	Password string       `json:"password"`
	GroupID  int          `json:"-"`
	Status   int          `json:"-"`
	Profile  AdminProfile `json:"profile"`

	service.Page `json:"page"`
}

type AdminProfile struct {
	AdminID uint   `json:"admin_id"`
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

func (a *Admin) Get() ([]Admin, error) {
	models, err := admin.Get(a.Page.Num, a.Page.Size, a.getMaps())
	if err != nil {
		return nil, err
	}

	admins := []Admin{}
	copier.Copy(&admins, &models)

	return admins, nil
}

func (a *Admin) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
