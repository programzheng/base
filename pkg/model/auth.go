package model

import "github.com/jinzhu/gorm"

type AdminLogin struct {
	gorm.Model
	AdminID  uint   `gorm:"not null"`
	Token    string `gorm:"not null"`
	Remember bool   `gorm:"default:false"`
	IP       string `gorm:"not null"`
}

func AddAdminLogin(adminLogin AdminLogin) error {
	Migrate(&adminLogin)
	if err := db.Create(&adminLogin).Error; err != nil {
		return err
	}

	return nil
}

func GetAdminLogin(adminLogin AdminLogin) (*AdminLogin, error) {
	if err := db.Where(&adminLogin).First(&adminLogin).Error; err != nil {
		return nil, err
	}
	return &adminLogin, nil
}
