package admin

import (
	"github.com/jinzhu/gorm"
)

//Login is vaildation struct
type Login struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Admin struct {
	gorm.Model
	Account  string `json:"account" gorm:"unique; not null"`
	Password string `json:"password" gorm:"unique; not null"`
	GroupID  int    `gorm:"index"`
	Status   int    `gorm:"defalut:0"`
	Profile  AdminProfile
}

type AdminLogin struct {
	gorm.Model
	AdminID uint   `gorm:"not null"`
	Token   string `gorm:"not null"`
	IP      string `gorm:"not null"`
}

type AdminProfile struct {
	gorm.Model
	AdminID uint
	Name    string `json:"name"`
}
