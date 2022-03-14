package email

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

type Email struct {
	FROM    string
	TO      []string
	BCC     []string
	CC      []string
	SUBJECT string
	HTML    []byte
}

func generateHtmlStringByHtmlFile(path string, data interface{}) string {
	tplString := ""

	var tmpl, err = template.ParseFiles(path)
	if err != nil {
		log.Printf("SendEmailByText template.Must error:%v", err)
	}

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, data); err != nil {
		if err != nil {
			log.Printf("SendEmailByText tmpl.Execute error:%v", err)
		}
	}
	tplString = tpl.String()

	return tplString
}

func (e *Email) SendEmailByHtml() error {
	newEmail := email.NewEmail()
	newEmail.From = e.FROM
	newEmail.To = e.TO
	newEmail.Bcc = e.BCC
	newEmail.Cc = e.CC
	newEmail.Subject = e.SUBJECT
	newEmail.HTML = e.HTML

	err := newEmail.Send(
		viper.GetString("NOTIFICATION_EMAIL_HOST")+":"+viper.GetString("NOTIFICATION_EMAIL_PORT"),
		smtp.PlainAuth("", viper.GetString("NOTIFICATION_EMAIL_USERNAME"),
			viper.GetString("NOTIFICATION_EMAIL_PASSWORD"),
			viper.GetString("NOTIFICATION_EMAIL_HOST"),
		),
	)

	return err
}
