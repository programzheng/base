package bot

import (
	"base/pkg/helper"
	"base/pkg/library/line/bot/template"
	"strings"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/line/line-bot-sdk-go/linebot"
	log "github.com/sirupsen/logrus"
)

func GroupParseTextGenTemplate(lineId LineID, text string) linebot.SendingMessage {
	parseText := strings.Split(text, "|")

	//功能說明
	if len(parseText) == 1 {
		switch parseText[0] {
		case "c helper", "記帳說明", "記帳":
			return template.Text("\"記帳\"\n將按照群組人數去做平均計算，使用記帳請使用以下格式輸入\n\"記帳|標題|總金額|備註\"\n例如:\n記帳|生日聚餐|1234|本人生日")
		case "c list helper", "記帳列表說明":
			return template.Text("\"記帳列表\"\n將回傳記帳紀錄的列表，格式為:\n日期時間 標題|金額| 平均金額 |付款人|備註")
		}
	}
	//功能
	switch parseText[0] {
	// c list||記帳列表
	case "c list", "記帳列表":
		lb := LineBilling{}
		where := make(map[string]interface{})
		not := make(map[string]interface{})
		not["group_id"] = ""
		lbs, err := lb.Get(where, not)
		if err != nil {
			log.Fatal("取得記帳列表錯誤", err)
		}
		if len(lbs) == 0 {
			return template.Text("沒有記帳紀錄哦!")
		}
		dstByUserID := make(map[string]string, 0)
		underscore.Chain(lbs).DistinctBy("UserID").SelectMany(func(lb LineBilling, _ int) map[string]string {
			lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
			if err != nil {
				log.Fatal("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
			}
			dst := make(map[string]string)
			dst[lb.UserID] = lineMember.DisplayName
			return dst
		}).Value(&dstByUserID)
		var sb strings.Builder
		sb.Grow(len(lbs))
		for _, lb := range lbs {
			memberName := "Unknow"
			//check line member display name is exist
			if _, ok := dstByUserID[lb.UserID]; ok {
				memberName = dstByUserID[lb.UserID]
			}
			amountAvg, amountAvgBase := calculateAmount(lineId.GroupID, helper.ConvertToFloat64(lb.Billing.Amount))
			text := lb.Billing.CreatedAt.Format(helper.Yyyymmddhhmmss) + " " +
				lb.Billing.Title + "|" + helper.ConvertToString(lb.Billing.Amount) + "/" + helper.ConvertToString(amountAvgBase) + " = " + helper.ConvertToString(amountAvg) + " |" + memberName + "|" + lb.Billing.Note + "\n"
			sb.WriteString(text)
		}
		//get user id total amount

		return template.Text(sb.String())
	// c||記帳|生日聚餐|1234|本人生日
	case "c", "記帳":
		title := parseText[1]
		amount := helper.ConvertToInt(parseText[2])
		note := parseText[3]
		billingAction(lineId, amount, title, note)
		amountFloat64 := helper.ConvertToFloat64(amount)
		amountAvg, amountAvgBase := calculateAmount(lineId.GroupID, amountFloat64)
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

func calculateAmount(groupID string, amount float64) (float64, int) {
	//預設平均計算基數
	amountAvgBase := 3.0
	groupMemberCount, err := botClient.GetGroupMemberCount(groupID).Do()
	if err != nil {
		log.Fatal("line messaging api get group member count error:", err)
	}
	amountAvgBase = helper.ConvertToFloat64(groupMemberCount.Count)
	amountAvg := amount / amountAvgBase
	return amountAvg, groupMemberCount.Count
}
