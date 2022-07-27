package email

import (
	"testing"

	"github.com/programzheng/base/config"
	_ "github.com/programzheng/base/config"
)

func TestSendEmailByText(t *testing.T) {
	data := struct {
		Key string
	}{
		Key: "key",
	}
	htmlString := generateHtmlStringByHtmlFile("dist/test.html", data)

	email := Email{
		FROM:    config.Cfg.GetString("NOTIFICATION_EMAIL_FROM"),
		TO:      []string{"test@example.com"},
		SUBJECT: "Test Email",
		HTML:    []byte(htmlString),
	}

	err := email.SendEmailByHtml()

	if err != nil {
		t.Errorf("TestSendEmailByText error:%v", err)
	}

	t.Log("success")
}
