package email

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type EmailConfig struct {
	FromEmail string
	ApiKey    string
	SMTPUser  string
	SMTPHost  string
	SMTPPort  int
}

func Send(config EmailConfig, toEmail string, subject string, content string) error {
	if config.FromEmail == "" || config.ApiKey == "" {
		return fmt.Errorf("email config: missing mandatory parameters")
	}

	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort
	if smtpHost == "" || smtpPort == 0 {
		smtpHost = "live.smtp.mailtrap.io"
		smtpPort = 587
	}

	message := gomail.NewMessage()

	message.SetHeader("From", config.FromEmail)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", content)

	dialer := gomail.NewDialer(smtpHost, smtpPort, config.SMTPUser, config.ApiKey)
	err    := dialer.DialAndSend(message)

	return err
}
