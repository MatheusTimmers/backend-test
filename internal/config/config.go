package config

import (
	"log"
	"os"
)

type AppConfig struct {
	EmailFrom     string
	EmailAPIKey   string
	SMTPHost      string
	SMTPPort      string
	SMTPUser      string
	AdminToken    string
}

var Config *AppConfig

func Load() {
	Config = &AppConfig{
		EmailFrom:     os.Getenv("EMAIL_API"),
		EmailAPIKey:   os.Getenv("EMAIL_API_KEY"),
		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		AdminToken:    os.Getenv("ADMIN_TOKEN"),
		SMTPUser:      os.Getenv("SMTP_USER"),
	}

	if Config.EmailFrom == "" || Config.EmailAPIKey == "" || Config.AdminToken == "" {
		log.Fatal("variáveis de ambiente obrigatórias não estão definidas")
	}
}
