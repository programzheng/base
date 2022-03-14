package auth

import (
	"github.com/programzheng/base/pkg/model"

	"gorm.io/gorm"
)

type AdminLogin struct {
	gorm.Model
	AdminID  uint   `gorm:"not null"`
	Token    string `gorm:"not null"`
	Remember bool   `gorm:"default:false"`
	IP       string `gorm:"not null"`
}

func AddAdminLogin(adminLogin AdminLogin) error {
	model.SetupTableModel(&adminLogin)
	if err := model.GetDB().Create(&adminLogin).Error; err != nil {
		return err
	}

	return nil
}

func GetAdminLogin(adminLogin AdminLogin) (*AdminLogin, error) {
	if err := model.GetDB().Where(&adminLogin).First(&adminLogin).Error; err != nil {
		return nil, err
	}
	return &adminLogin, nil
}
