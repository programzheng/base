package bot

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"base/pkg/helper"
	"base/pkg/job/line"
	"base/pkg/library/line/bot/template"
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

func ParseTextGenTemplate(toID string, text string) linebot.SendingMessage {
	parseText := strings.Split(text, "|")

	if len(parseText) == 1 {

	}
	switch parseText[0] {
	// 記帳|測試|300|備註
	case "記帳":
		title := parseText[1]
		amount := helper.ConvertToInt(parseText[2])
		note := parseText[4]
		billingAction(toID, title, amount, note)
		amountAvg := amount / 3
		return template.Text("記帳完成，" + parseText[2] + "/3=" + helper.ConvertToString(amountAvg))
	case "TODO":
		date := parseText[1]
		replyText := parseText[2]
		parseDate := strings.Split(date, " ")
		switch parseDate[0] {
		case "every":
			// TODO|every 19:55|測試29號13:30送出
			todoAction(toID, "every", parseDate[1], template.TODO(replyText))
			return template.Text("設置完成將於每天" + parseDate[1] + "\n傳送訊息:" + replyText)
		default:
			// TODO|2020/02/29 13:00|測試29號13:30送出
			todoAction(toID, "once", date, template.TODO(replyText))
			return template.Text("設置完成將於" + date + "\n傳送訊息:" + replyText)
		}

	}
	return template.Text(text)
}

func LineReplyMessage(replyToken string, messages ...linebot.SendingMessage) {
	_, err := botClient.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		log.Println("LINE Message API parse Request error:", err)
	}
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
		timeRange := helper.CalcTimeRange(time.Now().Format(helper.GetTimeLayout()), date)
		jobrunner.In(time.Duration(timeRange)*time.Second, job)
	}
}

func billingAction(toID string, title string, amount int, note string) (billing.Billing, bot.LineBilling) {
	b := billing.Billing{
		Title:  title,
		Amount: amount,
		Note:   note,
	}
	billing, err := b.Add()
	if err != nil {
		log.Fatal("billingAction add error:", err)
	}
	lb := bot.LineBilling{
		BillingID: billing.ID,
		UserID:    toID,
	}
	return b, lb
}
