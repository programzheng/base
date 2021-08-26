package bot

import (
	"base/pkg/helper"
	"base/pkg/library/line/bot/template"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	log "github.com/sirupsen/logrus"
)

func UserParseTextGenTemplate(lineId LineID, text string) linebot.SendingMessage {
	parseText := strings.Split(text, "|")

	if len(parseText) == 1 {

	}
	switch parseText[0] {
	// 記帳列表
	case "記帳列表":
		lb := LineBilling{}
		where := make(map[string]interface{})
		not := make(map[string]interface{})
		lbs, err := lb.Get(where, not)
		if err != nil {
			log.Fatal("取得記帳列表錯誤", err)
		}
		var sb strings.Builder
		sb.Grow(len(lbs))
		for _, lb := range lbs {
			memberName := "Unknow"
			lineMember, err := botClient.GetProfile(lb.UserID).Do()
			if err != nil {
				log.Fatal("line messaging api get member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
			}
			memberName = lineMember.DisplayName
			text := lb.Billing.CreatedAt.Format(helper.Yyyymmddhhmmss) + " " +
				lb.Billing.Title + "|" + helper.ConvertToString(lb.Billing.Amount) + "|" + lb.Billing.Note + "|" + memberName + "\n"
			sb.WriteString(text)
		}
		return template.Text(sb.String())
	// 記帳|測試|300|備註
	case "記帳":
		title := parseText[1]
		amount := helper.ConvertToInt(parseText[2])
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
