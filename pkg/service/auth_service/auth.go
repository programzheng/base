package auth_service

import (
	"github.com/ProgramZheng/base/pkg/model"
)

//Login is vaildation struct
type Login struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminLogin struct {
	AdminID  uint
	Token    string
	Remember bool
	IP       string
}

func (al *AdminLogin) AddAdminLogin() error {
	adminLogin := model.AdminLogin{
		AdminID:  al.AdminID,
		Token:    al.Token,
		Remember: al.Remember,
		IP:       al.IP,
	}
	if err := model.AddAdminLogin(adminLogin); err != nil {
		return err
	}
	return nil
}

func (al *AdminLogin) GetAdminLogin() (*model.AdminLogin, error) {
	where := model.AdminLogin{
		AdminID: al.AdminID,
		Token:   al.Token,
		IP:      al.IP,
	}
	adminLogin, err := model.GetAdminLogin(where)
	if err != nil {
		return nil, err
	}
	return adminLogin, nil
}

// func (al *AddAdminLogin) Edit() (*model.AddAdminLogin) {

// }
