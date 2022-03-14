package email

import (
	"testing"

	_ "github.com/programzheng/base/config"
	"github.com/spf13/viper"
)

func TestSendEmailByText(t *testing.T) {
	data := struct {
		Key string
	}{
		Key: "key",
	}
	htmlString := generateHtmlStringByHtmlFile("dist/test.html", data)

	email := Email{
		FROM:    viper.GetString("NOTIFICATION_EMAIL_FROM"),
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
