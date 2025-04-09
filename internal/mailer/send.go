package mailer

import (
	"github.com/MatheusTimmers/backend-test/internal/logger"
	"github.com/MatheusTimmers/backend-test/pkg/email"
	"github.com/MatheusTimmers/backend-test/pkg/models"
)

func SendEmail(user models.User, emailType EmailType) {
	go sendEmail(user, emailType)
}

func sendEmail(user models.User, t EmailType) {
	subject, body := getEmailContent(t, user)

	err := email.Send(user.Name, user.Email, subject, body)
	if err != nil {
		logger.Log.Error(err.Error())
	}
}
