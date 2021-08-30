package bot

import (
	"base/pkg/library/line/bot/template"
	"strings"
)

func UserParseTextGenTemplate(lineId LineID, text string) interface{} {
	parseText := strings.Split(text, "|")

	if len(parseText) == 1 {

	}
	switch parseText[0] {
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
