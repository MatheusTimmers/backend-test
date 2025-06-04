package usecase

type EmailType int

const (
	NewUser EmailType = iota
	NewPoint
	Winner
)

type Mailer interface {
	SendEmail(name, toEmail, inviteCode string, emailType EmailType) error
}
