package mailer

import (
	"github.com/MatheusTimmers/backend-test/internal/config"
	"github.com/MatheusTimmers/backend-test/internal/logger"
	"github.com/MatheusTimmers/backend-test/pkg/email"
	"github.com/MatheusTimmers/backend-test/pkg/models"
)

func SendEmail(user models.User, emailType EmailType) {
	go sendEmail(user, emailType)
}

func sendEmail(user models.User, t EmailType) {
	subject, body := getEmailContent(t, user)

	c := email.EmailConfig{
		FromEmail: config.Config.EmailFrom,
		ApiKey:    config.Config.EmailAPIKey,
		SMTPUser:  config.Config.SMTPUser,
		SMTPHost:  config.Config.SMTPHost,
		SMTPPort:  config.Config.SMTPPort,
	}

	err := email.Send(c, user.Email, subject, body)
	if err != nil {
		logger.Log.Errorf("error sending email to %s: %v", user.Email, err)
	} else {
		logger.Log.Infof("email sent successfully to %s", user.Email)
	}
}
