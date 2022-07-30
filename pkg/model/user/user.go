package user

import (
	"github.com/programzheng/base/pkg/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `gorm:"unique; not null"`
	Account  string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null" json:"-"`
	Profile  UserProfile
}

type UserProfile struct {
	gorm.Model
	UserID uint
	Name   string
}

func (u *User) Add() (*User, error) {
	if err := model.GetDB().Create(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) First() (*User, error) {
	if err := model.GetDB().First(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
