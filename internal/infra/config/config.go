package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Application *application
	Database    *database
	Mailer      *mailer
}

type application struct {
	AdminToken string
	ServerPort string
	ServerHost string
}

type database struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SslMode  string
	LogLevel string
}

type mailer struct {
	FromEmail    string
	FromPassword string
	SMTPHost     string
	SMTPPort     int
}

func Config() *AppConfig {
	return &AppConfig{
		Application: &application{
			AdminToken: os.Getenv("ADMIN_TOKEN"),
			ServerPort: os.Getenv("SERVER_PORT"),
			ServerHost: os.Getenv("SERVER_HOST"),
		},
		Database: &database{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SslMode:  os.Getenv("DB_SSLMODE"),
			LogLevel: os.Getenv("DB_LOG_LEVEL"),
		},
		Mailer: &mailer{
			FromEmail:    os.Getenv("FROM_EMAIL"),
			FromPassword: os.Getenv("FROM_PASSWORD"),
			SMTPHost:     os.Getenv("SMTP_HOST"),
			SMTPPort:     parseInt(os.Getenv("SMTP_PORT"), 587),
		},
	}
}

func parseInt(s string, defaultVal int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return i
}
