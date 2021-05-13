package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Telegram holds configuration related to Telegram bot.
type Telegram struct {
	RecipientID int    `env:"TELEGRAM_RECIPIENT_ID,required"`
	URL         string `env:"TELEGRAM_URL,required"`
	Token       string `env:"TELEGRAM_TOKEN,required"`
}

// TaniFund holds configuration related to TaniFund.
type TaniFund struct {
	URL string `env:"TANIFUND_URL,required"`
}

// Config holds configuration for the project.
type Config struct {
	Telegram Telegram
	TaniFund TaniFund
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
