package util

import (
	"net/smtp"
	"log"
)

func Send(to string, subject string, body string) (error, string) {
	from := "development.xx@gmail.com"
	pass := "your_password"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: "+subject+"\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err, "Failed"
	} else {
		log.Print("message sent to "+ to)
		return nil, "Ok"
	}

}
