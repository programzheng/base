package bot

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/service/bot"
	"github.com/spf13/viper"
)

func setLineBot() *linebot.Client {
	channelSecret := viper.Get("LINE_CHANNEL_SECRET").(string)
	channelAccessToken := viper.Get("LINE_CHANNEL_ACCESS_TOKEN").(string)
	botClient, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Println("LINE bot error:", err)
	}
	return botClient
}

func LineWebHook(ctx *gin.Context) {
	botClient := setLineBot()
	events, err := botClient.ParseRequest(ctx.Request)
	if err != nil {
		log.Println("LINE Message API parse Request error:", err)
	}

	for _, event := range events {
		fmt.Println(event.Source.UserID)
		request, err := event.MarshalJSON()
		if err != nil {
			log.Println("LINE Message API event to json error:", err)
		}
		requestString := string(request)
		fmt.Println(requestString)
		botService := bot.LineBotRequest{
			Type:       string(event.Source.Type),
			GroupID:    event.Source.GroupID,
			RoomID:     event.Source.RoomID,
			UserID:     event.Source.UserID,
			ReplyToken: event.ReplyToken,
			Request:    requestString,
		}
		if _, err := botService.Add(); err != nil {
			function.Fail(ctx, err)
			return
		}
		// function.GetType(event.Message)
		// switch t := event.Message.(type) {
		// case *linebot.TextMessage:
		// 	fmt.Println(event.Message)
		// }
		// message := defaultMessage(event.Message())
		message := defaultTemplateMessage()
		lineReplyMessage(botClient, event.ReplyToken, message)
	}
}

func lineReplyMessage(botClient *linebot.Client, replyToken string, message *linebot.TemplateMessage) {
	var messages []linebot.SendingMessage
	messages = append(messages, message)
	_, err := botClient.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		log.Println("LINE Message API parse Request error:", err)
	}
}

func defaultMessage(text string) *linebot.TextMessage {
	message := linebot.NewTextMessage(text)
	return message
}

func defaultTemplateMessage() *linebot.TemplateMessage {
	leftBtn := linebot.NewMessageAction("left", "left clicked")
	rightBtn := linebot.NewMessageAction("right", "right clicked")
	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
	message := linebot.NewTemplateMessage("Reply", template)
	return message
}
