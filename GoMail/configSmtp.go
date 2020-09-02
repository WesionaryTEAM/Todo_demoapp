package gomail

import (
	"fmt"
	"net/smtp"
	"os"
)

var emailAuth smtp.Auth

//configSMTPSendEmail is ...
func configSMTPSendEmail(to []string, data interface{}, template string) bool {
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPort := os.Getenv("EMAIL_PORT")
	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody := parseTemplate(template, data)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false
	}
	return true

}
