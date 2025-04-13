package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Application *application
	Database *database
	Mailer *mailer
}

type application struct {
	AdminToken    string
	ServerPort    string
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
	EmailFrom     string
	EmailAPIKey   string
	SMTPHost      string
	SMTPPort      int
	SMTPUser      string
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
			EmailFrom:   os.Getenv("EMAIL_API"),
			EmailAPIKey: os.Getenv("EMAIL_API_KEY"),
			SMTPHost:    os.Getenv("SMTP_HOST"),
			SMTPPort:    parseInt(os.Getenv("SMTP_PORT"), 587),
			SMTPUser:    os.Getenv("SMTP_USER"),
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
