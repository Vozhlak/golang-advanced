package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db     DbConfig
	Auth   AuthConfig
	Verify VerifyConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type VerifyConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default config", err)
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Verify: VerifyConfig{
			Email:    os.Getenv("VERIFY_EMAIL"),
			Password: os.Getenv("VERIFY_PASSWORD"),
			Address:  os.Getenv("VERIFY_ADDRESS"),
		},
	}
}
