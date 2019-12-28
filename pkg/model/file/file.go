package file

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/model"
)

type File struct {
	gorm.Model
	System string
	Type   string
	Path   string
	Name   string
}

func Add(file File) error {
	model.Migrate(&file)
	if err := model.DB.Save(&file).Error; err != nil {
		return err
	}

	return nil
}
