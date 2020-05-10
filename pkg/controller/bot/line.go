package bot

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service/bot"
	log "github.com/sirupsen/logrus"
)

const lineOfficalID = "Udeadbeefdeadbeefdeadbeefdeadbeef"

var botClient = bot.SetLineBot()

func LineWebHook(ctx *gin.Context) {
	events, err := botClient.ParseRequest(ctx.Request)
	if err != nil {
		log.Println("LINE Message API parse Request error:", err)
	}

	for _, event := range events {
		request, err := event.MarshalJSON()
		if err != nil {
			log.Println("LINE Message API event to json error:", err)
		}
		if event.Source.UserID == lineOfficalID {
			helper.Success(ctx, nil, nil)
			return
		}
		requestString := string(request)
		lineBotRequest := bot.LineBotRequest{
			Type:       string(event.Source.Type),
			GroupID:    event.Source.GroupID,
			RoomID:     event.Source.RoomID,
			UserID:     event.Source.UserID,
			ReplyToken: event.ReplyToken,
			Request:    requestString,
		}
		if _, err := lineBotRequest.Add(); err != nil {
			helper.Fail(ctx, err)
			return
		}
		switch event.Source.Type {
		case "user":
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyTemplateMessage := bot.ParseTextGenTemplate(event.Source.UserID, message.Text)
					bot.LineReplyMessage(event.ReplyToken, replyTemplateMessage)
				}
			}
		case "group":
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyTemplateMessage := bot.ParseTextGenTemplate(event.Source.GroupID, message.Text)
					bot.LineReplyMessage(event.ReplyToken, replyTemplateMessage)
				}
			}
		}

	}
}

func defaultTemplateMessage() *linebot.TemplateMessage {
	leftBtn := linebot.NewMessageAction("left", "left clicked")
	rightBtn := linebot.NewMessageAction("right", "right clicked")
	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
	message := linebot.NewTemplateMessage("Reply", template)
	return message
}
