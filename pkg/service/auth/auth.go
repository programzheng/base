package auth

import (
	"base/pkg/model/auth"
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
	adminLogin := auth.AdminLogin{
		AdminID:  al.AdminID,
		Token:    al.Token,
		Remember: al.Remember,
		IP:       al.IP,
	}
	if err := auth.AddAdminLogin(adminLogin); err != nil {
		return err
	}
	return nil
}

func (al *AdminLogin) GetAdminLogin() (*auth.AdminLogin, error) {
	where := auth.AdminLogin{
		AdminID: al.AdminID,
		Token:   al.Token,
		IP:      al.IP,
	}
	adminLogin, err := auth.GetAdminLogin(where)
	if err != nil {
		return nil, err
	}
	return adminLogin, nil
}

// func (al *AddAdminLogin) Edit() (*auth.AddAdminLogin) {

// }
