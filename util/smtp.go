package util

import (
	"net/smtp"
	"log"
)
// Sometime need to change security setting on your google account setting :
// Connected apps & sites -> Allow less secure apps (ON)
func Send(to string, subject string, body string, smtpUser string, smtpPassw string,
	  smtpHost string, smtpAddr string ) (error, string) {
	from := smtpUser
	pass := smtpPassw
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: "+subject+"\n\n" +
		body

	err := smtp.SendMail(smtpAddr,
		smtp.PlainAuth("", from, pass, smtpHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err, "Failed"
	} else {
		log.Print("message sent to "+ to)
		return nil, "Ok"
	}

}
