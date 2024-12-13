package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config  mmenghubungkan database Postgresql "go_education"
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// TwilioConfig menghubungkan API dari WhatsApp
type TwilioConfig struct {
	AccountSID     string
	AuthToken      string
	WhatsAppNumber string
}

// LoadTwilioConfig memanggil isi dari file .env
var Twilio TwilioConfig

func LoadTwilioConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Twilio = TwilioConfig{
		AccountSID:     os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:      os.Getenv("TWILIO_AUTH_TOKEN"),
		WhatsAppNumber: os.Getenv("TWILIO_WHATSAPP_NUMBER"),
	}
}

// LoadConfig memanggil isi dari file .env
func LoadConfig() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
