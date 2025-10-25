package env

import "github.com/caarlos0/env/v11"

type botEnvConfig struct {
	Token string `env:"TELEGRAM_BOT_TOKEN,required"`
}

type botConfig struct {
	row botEnvConfig
}

func NewBotConfig() (*botConfig, error) {
	var row botEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &botConfig{row: row}, nil
}

func (b *botConfig) Token() string {
	return b.row.Token
}
