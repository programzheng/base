package bot

import (
	"github.com/programzheng/base/pkg/model/bot"
)

type LineBotRequest struct {
	Type       string
	GroupID    string
	RoomID     string
	UserID     string
	ReplyToken string
	Request    string
}

func (lineBotRequest *LineBotRequest) Add() (uint, error) {
	model := bot.LineBotRequest{
		Type:       lineBotRequest.Type,
		GroupID:    lineBotRequest.GroupID,
		RoomID:     lineBotRequest.RoomID,
		UserID:     lineBotRequest.UserID,
		ReplyToken: lineBotRequest.ReplyToken,
		Request:    lineBotRequest.Request,
	}
	ID, err := bot.Add(model)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
