package auth

import (
	"github.com/programzheng/base/pkg/model"

	"github.com/jinzhu/gorm"
)

type AdminLogin struct {
	gorm.Model
	AdminID  uint   `gorm:"not null"`
	Token    string `gorm:"not null"`
	Remember bool   `gorm:"default:false"`
	IP       string `gorm:"not null"`
}

func AddAdminLogin(adminLogin AdminLogin) error {
	model.Migrate(&adminLogin)
	if err := model.DB.Create(&adminLogin).Error; err != nil {
		return err
	}

	return nil
}

func GetAdminLogin(adminLogin AdminLogin) (*AdminLogin, error) {
	if err := model.DB.Where(&adminLogin).First(&adminLogin).Error; err != nil {
		return nil, err
	}
	return &adminLogin, nil
}
