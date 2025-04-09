package mailer

import (
	"fmt"

	"github.com/MatheusTimmers/backend-test/pkg/models"
)

type EmailType int

const (
	NewPoint EmailType = iota
	NewUser
	Winner
)

func getEmailContent(t EmailType, user models.User) (string, string) {
	switch t {
	case NewPoint:
		return "Você ganhou um ponto!",
			fmt.Sprintf(`
			<h2>Parabéns, %s!</h2>
			<p>Alguém usou seu código de convite e você ganhou um ponto!</p>
		`, user.Name)

	case Winner:
		return "Você está no Top 10!",
			fmt.Sprintf(`
			<h2>Parabéns, %s!</h2>
			<p>Você ficou entre os 10 primeiros colocados!</p>
		`, user.Name)

	case NewUser:
		return "Bem-vindo à nossa competição",
			fmt.Sprintf(`
			<h1>Olá, %s!</h1>
			<p>Obrigado por participar da nossa competição!</p>
			<p>Compartilhe seu link <a href="http://localhost:8080/register?invite=%s">de convite</a> e ganhe pontos!</p>
		`, user.Name, user.InviteCode)

	default:
		return "Mensagem", "Tipo de e-mail não reconhecido."
	}
}
