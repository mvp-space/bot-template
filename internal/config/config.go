// nolint
package config

import (
	"gopkg.in/yaml.v3"
	"os"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-env"

	"github.com/mvp-space/bot-template/pkg/log"
)

// Config represents an application configuration.
type Config struct {
	// Application environment. required.
	IsDebug bool `yaml:"is_debug" env:"IS_DEBUG"`
	// Telegram token. required.
	TelegramToken string `yaml:"telegram_token" env:"TELEGRAM_TOKEN,secret"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IsDebug, validation.Required),
		validation.Field(&c.TelegramToken, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{}

	// load from YAML config file
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
