package scratches

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	WhatsAppAccessToken string
	WhatsAppPhoneID     string
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleRedirectURI   string
}

// LoadConfig loads configuration from environment variables
// It returns an error if required variables are missing
func LoadConfig() (*Config, error) {
	// Try to load .env file, but don't fail if it doesn't exist
	// This allows for flexibility in different environments
	if err := godotenv.Load(); err != nil {
		// Log warning but continue - env vars might be set another way
		fmt.Printf("Warning: could not load .env file: %v\n", err)
	}

	cfg := &Config{
		WhatsAppAccessToken: strings.TrimSpace(os.Getenv("WHATSAPP_ACCESS_TOKEN")),
		WhatsAppPhoneID:     strings.TrimSpace(os.Getenv("WHATSAPP_PHONE_NUMBER_ID")),
		GoogleClientID:      strings.TrimSpace(os.Getenv("GOOGLE_OAUTH_CLIENT_ID")),
		GoogleClientSecret:  strings.TrimSpace(os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")),
		GoogleRedirectURI:   strings.TrimSpace(os.Getenv("GOOGLE_OAUTH_REDIRECT_URI")),
	}

	// Validate required fields
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// Validate checks if all required configuration values are present
func (c *Config) Validate() error {
	var missing []string

	if c.WhatsAppAccessToken == "" {
		missing = append(missing, "WHATSAPP_ACCESS_TOKEN")
	}
	if c.WhatsAppPhoneID == "" {
		missing = append(missing, "WHATSAPP_PHONE_NUMBER_ID")
	}
	if c.GoogleClientID == "" {
		missing = append(missing, "GOOGLE_OAUTH_CLIENT_ID")
	}
	if c.GoogleClientSecret == "" {
		missing = append(missing, "GOOGLE_OAUTH_CLIENT_SECRET")
	}
	if c.GoogleRedirectURI == "" {
		missing = append(missing, "GOOGLE_OAUTH_REDIRECT_URI")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s",
			strings.Join(missing, ", "))
	}

	return nil
}

// LoadConfigWithDefaults loads config with optional default values
func LoadConfigWithDefaults(defaults *Config) (*Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// Apply defaults for empty values
	if defaults != nil {
		if cfg.GoogleRedirectURI == "" && defaults.GoogleRedirectURI != "" {
			cfg.GoogleRedirectURI = defaults.GoogleRedirectURI
		}
		// Add other defaults as needed
	}

	return cfg, nil
}

// MustLoadConfig loads config and panics on error
// Use this only in main() or init() functions where you want to fail fast
func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}
