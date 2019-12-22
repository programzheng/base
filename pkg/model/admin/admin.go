package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/model"
)

type Admin struct {
	gorm.Model
	Account  string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null" json:"-"`
	Status   int    `gorm:"defalut:0"`
	Profile  AdminProfile
}

type AdminProfile struct {
	gorm.Model
	AdminID uint
	Name    string
}

func Add(admin Admin) error {
	model.Migrate(&admin, &admin.Profile)
	if err := model.DB.Save(&admin).Error; err != nil {
		return err
	}

	return nil
}

func GetForLogin(admin Admin) (*Admin, error) {
	if err := model.DB.Where(&admin).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Admin, error) {
	var models []*Admin
	err := model.DB.Preload("Profile").Where(maps).Offset(pageNum).Limit(pageSize).Find(&models).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return models, nil
}
