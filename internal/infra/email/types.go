package email

import (
	"bytes"
	"html/template"

	"github.com/MatheusTimmers/backend-test/internal/application/usecase"
	"github.com/MatheusTimmers/backend-test/pkg/logger"
)

type templateData struct {
	Name       string
	InviteCode string
}

func getEmailContent(t usecase.EmailType, name, inviteCode string) (string, string) {
	var tplName string
	var subject string

	switch t {
	case usecase.NewPoint:
		subject = "Você ganhou um ponto!"
		tplName = "new_point.html"
	case usecase.Winner:
		subject = "Você está no Top 10!"
		tplName = "winner.html"
	case usecase.NewUser:
		subject = "Bem-vindo à nossa competição"
		tplName = "new_user.html"
	default:
		return "Mensagem", "Tipo de e-mail não reconhecido."
	}

	tmpl, err := template.ParseFiles("internal/infra/email/templates/" + tplName)
	if err != nil {
		logger.Log().Error("erro to parse template", err)
		return subject, "Erro ao carregar template."
	}

	data := templateData{Name: name, InviteCode: inviteCode}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return subject, "Erro ao processar template."
	}

	return subject, buf.String()
}
