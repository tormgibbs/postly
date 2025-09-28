package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WhatsAppAccessToken string
	WhatsAppPhoneID     string
	WhatsAppWebhookToken string

	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
}

func LoadConfig() *Config {
	godotenv.Load()

	cfg := &Config{
		WhatsAppAccessToken: os.Getenv("WHATSAPP_ACCESS_TOKEN"),
		WhatsAppPhoneID:     os.Getenv("WHATSAPP_PHONE_NUMBER_ID"),
		WhatsAppWebhookToken: os.Getenv("WHATSAPP_WEBHOOK_VERIFY_TOKEN"),
		GoogleClientID:      os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		GoogleClientSecret:  os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		GoogleRedirectURI:   os.Getenv("GOOGLE_OAUTH_REDIRECT_URI"),
	}

	return cfg
}
