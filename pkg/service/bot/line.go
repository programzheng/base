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

type LineID struct {
	GroupID string
	RoomID  string
	UserID  string
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

func ParseTextGenTemplate(lineId LineID, text string) linebot.SendingMessage {
	parseText := strings.Split(text, "|")

	if len(parseText) == 1 {

	}
	switch parseText[0] {
	// 記帳列表
	case "記帳列表":

	// 記帳|測試|300|備註
	case "記帳":
		amount := helper.ConvertToInt(parseText[1])
		title := parseText[2]
		note := parseText[3]
		billingAction(lineId, amount, title, note)
		amountFloat64 := helper.ConvertToFloat64(amount)
		amountAvgBase := helper.ConvertToFloat64(3)
		amountAvg := amountFloat64 / amountAvgBase
		return template.Text(title + ":記帳完成," + parseText[2] + "/" + helper.ConvertToString(int(amountAvgBase)) + " = " + "*" + helper.ConvertToString(amountAvg) + "*")
	case "TODO":
		date := parseText[1]
		replyText := parseText[2]
		parseDate := strings.Split(date, " ")
		switch parseDate[0] {
		case "every":
			// TODO|every 19:55|測試29號13:30送出
			todoAction(lineId.UserID, "every", parseDate[1], template.TODO(replyText))
			return template.Text("設置完成將於每天" + parseDate[1] + "\n傳送訊息:" + replyText)
		default:
			// TODO|2020/02/29 13:00|測試29號13:30送出
			todoAction(lineId.UserID, "once", date, template.TODO(replyText))
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
