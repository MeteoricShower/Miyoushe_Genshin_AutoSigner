package email

import (
	"Miyoushe_Genshin_AutoSigner/config"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"mime"
	"time"
)

func SendEmail(toEmail, subject, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "op")+"<"+config.EmailUsername+">")
	m.SetHeader("To", toEmail, toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf("%s<br/>处理时间:%s", msg, time.Now().Local().Format("2006-01-02 15:04:05")))

	d := gomail.NewDialer(config.EmailHost, config.EmailPort, config.EmailUsername, config.EmailPass)

	if err := d.DialAndSend(m); err != nil {
		log.Println("通知邮件发送失败,err:", err.Error())
	}

}
