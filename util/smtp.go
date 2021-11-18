package util

import (
	"log"
	"net/smtp"
	"rabbitmq-demo/configs"
)

func Send(to string, subject string, body string, smtpUser string) error {
	from := smtpUser
	//pass := smtpPassw
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(configs.Config.SmtpAddress, nil,
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	} else {
		log.Print("message sent to " + to)
		return nil
	}

}
