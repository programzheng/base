package bot

import (
	"reflect"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"base/pkg/helper"
	"base/pkg/job/line"
	"base/pkg/model/bot"
	"base/pkg/service/billing"

	"github.com/bamzi/jobrunner"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

type LineBotRequest struct {
	Type       string
	GroupID    string
	RoomID     string
	UserID     string
	ReplyToken string
	Request    string
}

type LineID struct {
	GroupID string
	RoomID  string
	UserID  string
}

type LinePostBackAction struct {
	Action string
	Data   map[string]interface{}
}

type LineBotPushMessage struct {
	Token string `json:"token"`
	Text  string `json:"text"`
}

var botClient = SetLineBot()

func SetLineBot() *linebot.Client {
	channelSecret := viper.Get("LINE_CHANNEL_SECRET").(string)
	channelAccessToken := viper.Get("LINE_CHANNEL_ACCESS_TOKEN").(string)
	botClient, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Println("LINE bot error:", err)
	}
	return botClient
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
	ID, err := model.Add()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func LineReplyMessage(replyToken string, messages interface{}) {
	var sendMessages []linebot.SendingMessage
	rv := reflect.ValueOf(messages)
	if rv.Kind() == reflect.Slice {
		sendMessages = messages.([]linebot.SendingMessage)
	} else {
		sendMessages = append(sendMessages, messages.(linebot.SendingMessage))
	}
	basicResponse, err := botClient.ReplyMessage(replyToken, sendMessages...).Do()
	if err != nil {
		log.Println("LINE Message API reply message Request error:", err)
	}
	log.Printf("LINE Message API reply message Request response:%v\n", basicResponse)
}

func LinePushMessage(toID string, messages ...linebot.SendingMessage) {
	botClient.PushMessage(toID, messages...).Do()
}

func todoAction(toID string, cycle string, date string, template *linebot.TextMessage) {
	job := line.Todo{
		BotClient: botClient,
		ToID:      toID,
		Template:  template,
	}
	switch cycle {
	case "every":
		parseTime := strings.Split(date, ":")
		hour := parseTime[0]
		minute := parseTime[1]
		jobrunner.Schedule(minute+" "+hour+" * * *", job)
	default:
		timeRange := helper.CalcTimeRange(time.Now().Format(helper.Yyyymmddhhmmss), date)
		jobrunner.In(time.Duration(timeRange)*time.Second, job)
	}
}

func billingAction(lineId LineID, amount int, title string, note string) (billing.Billing, bot.LineBilling) {
	b := billing.Billing{
		Title:  title,
		Amount: amount,
		Note:   note,
	}
	billing, err := b.Add()
	if err != nil {
		log.Fatal("billingAction Billing add error:", err)
	}
	lb := bot.LineBilling{
		BillingID: billing.ID,
		GroupID:   lineId.GroupID,
		RoomID:    lineId.RoomID,
		UserID:    lineId.UserID,
	}
	_, err = lb.Add()
	if err != nil {
		log.Fatal("billingAction LineBilling add error:", err)
	}
	return b, lb
}
