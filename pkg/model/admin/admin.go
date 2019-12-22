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

func AddAdmin(admin Admin) error {
	model.Migrate(&admin, &admin.Profile)
	if err := model.DB.Save(&admin).Error; err != nil {
		return err
	}

	return nil
}

func GetAdmin(admin Admin) (*Admin, error) {
	if err := model.DB.Where(&admin).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func GetAdmins(pageNum int, pageSize int, maps interface{}) ([]*Admin, error) {
	var admins []*Admin
	err := model.DB.Preload("Profile").Where(maps).Offset(pageNum).Limit(pageSize).Find(&admins).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return admins, nil
}
