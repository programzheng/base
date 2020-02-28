package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Todo struct {
	BotClient *linebot.Client
	ToID      string
	Template  *linebot.TextMessage
}

func (todo Todo) Run() {
	todo.BotClient.PushMessage(todo.ToID, todo.Template).Do()
}
