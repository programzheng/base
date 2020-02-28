package template

import "github.com/line/line-bot-sdk-go/linebot"

func Text(text string) *linebot.TextMessage {
	message := linebot.NewTextMessage(text)
	return message
}

func TODO(text string) *linebot.TextMessage {
	message := linebot.NewTextMessage(text)
	return message
}
