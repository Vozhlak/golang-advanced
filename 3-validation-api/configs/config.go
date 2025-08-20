package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
	Auth   AuthConfig
	Email  EmailConfig
}

type ServerConfig struct {
	BaseUrl string
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailConfig struct {
	Address  string
	Password string
	SMTPHost string
	SMTPPort string
	From     string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default config", err)
	}

	return &Config{
		Server: ServerConfig{
			BaseUrl: os.Getenv("SERVER_BASE_URL"),
		},
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Email: EmailConfig{
			Address:  os.Getenv("EMAIL_ADDRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			SMTPHost: os.Getenv("EMAIL_SMTP_HOST"),
			SMTPPort: os.Getenv("EMAIL_SMTP_PORT"),
			From:     os.Getenv("EMAIL_FROM"),
		},
	}
}
