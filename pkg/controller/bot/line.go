package bot

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/library/line/bot/template"
	"github.com/programzheng/base/pkg/service/bot"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

var pushMessage bot.LineBotPushMessage

func LinePush(ctx *gin.Context) {
	if err := ctx.BindJSON(&pushMessage); err != nil {
		helper.BadRequest(ctx, err)
		return
	}
	token := helper.CreateMD5(time.Now().Format(helper.GetIso8601()))
	if pushMessage.Token != token {
		helper.Unauthorized(ctx, nil)
		return
	}
	bot.LinePushMessage(viper.Get("LINE_DEFAULT_PUSH_ID").(string), template.Text(pushMessage.Text))
}

func defaultTemplateMessage() *linebot.TemplateMessage {
	leftBtn := linebot.NewMessageAction("left", "left clicked")
	rightBtn := linebot.NewMessageAction("right", "right clicked")
	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
	message := linebot.NewTemplateMessage("Reply", template)
	return message
}
