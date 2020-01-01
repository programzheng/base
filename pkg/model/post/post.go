package post

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/model"
)

type Post struct {
	gorm.Model
	Image   string `sql:"type:json`
	Title   string
	Summary string
	Detail  string `sql:"type:text"`
}

func Add(post Post) error {
	model.Migrate(&post)
	if err := model.DB.Save(&post).Error; err != nil {
		return err
	}

	return nil
}

func Get(pageNum int, pageSize int, maps interface{}) ([]*Post, error) {
	var posts []*Post
	err := model.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return posts, nil
}
