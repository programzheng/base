package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/model/admin"
)

func CreateAdmin(db *gorm.DB, account string, password string) error {
	password = helper.CreateHash(password)
	return db.Create(&admin.Admin{Account: account, Password: password}).Error
}
