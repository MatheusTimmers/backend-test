package email

import (
	"fmt"
	"os"

	"github.com/MatheusTimmers/backend-test/internal/logger"

	gomail "gopkg.in/mail.v2"
)

func Send(toName string, toEmail string, subject string, content string) error {
	fromEmail := os.Getenv("EMAIL_API")
	fromName := os.Getenv("NAME_API_EMAIL")
	apiKey := os.Getenv("EMAIL_API_KEY")

	if fromEmail == "" || fromName == "" || apiKey == "" {
		return fmt.Errorf("email config: missing environment variables")
	}

	message := gomail.NewMessage()

	message.SetHeader("From", fromEmail)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", content)

	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", apiKey)
	err := dialer.DialAndSend(message)
	if err != nil {
		logger.Log.Errorf("error sending email to %s: %v", toEmail, err)
	} else {
		logger.Log.Infof("email sent successfully to %s", toEmail)
	}

	return err
}
