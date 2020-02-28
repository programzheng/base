package bot

import (
	"github.com/jinzhu/gorm"
	"github.com/programzheng/base/pkg/model"
)

type LineBotRequest struct {
	gorm.Model
	Type       string `gorm:"not null"`
	GroupID    string
	RoomID     string
	UserID     string `gorm:"not null"`
	ReplyToken string `gorm:"not null"`
	Request    string `sql:"type:text" gorm:"not null"`
}

func (lineBotRequest LineBotRequest) Add() (uint, error) {
	model.Migrate(&lineBotRequest)
	if err := model.DB.Save(&lineBotRequest).Error; err != nil {
		return 0, err
	}
	return lineBotRequest.ID, nil
}
