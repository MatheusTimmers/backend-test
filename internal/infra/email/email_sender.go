package email

import (
	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	"github.com/MatheusTimmers/backend-test/internal/infra/config"
	"github.com/MatheusTimmers/backend-test/pkg/email"
	"github.com/MatheusTimmers/backend-test/pkg/logger"
)

type SmtpMailer struct{}

func NewMailer() usecase.Mailer {
	return &SmtpMailer{}
}

func (m *SmtpMailer) SendEmail(name, emailAddr, inviteCode string, t usecase.EmailType) error {
	subject, body := getEmailContent(t, name, inviteCode)

	c := email.EmailConfig{
		FromEmail: config.Config().Mailer.EmailFrom,
		ApiKey:    config.Config().Mailer.EmailAPIKey,
		SMTPUser:  config.Config().Mailer.SMTPUser,
		SMTPHost:  config.Config().Mailer.SMTPHost,
		SMTPPort:  config.Config().Mailer.SMTPPort,
	}

	err := email.Send(c, emailAddr, subject, body)
	if err != nil {
		logger.Log.Errorf("error sending email to %s: %v", emailAddr, err)
		return err
	}

	logger.Log.Infof("email sent successfully to %s", emailAddr)
	return nil
}
