package env

import (
	"github.com/caarlos0/env/v11"
)

type redisEnvConfig struct {
	Host              string `env:"REDIS_HOST,required"`
	Port              string `env:"REDIS_PORT,required"`
	ConnectionTimeout string `env:"REDIS_CONNECTION_TIMEOUT,required"`
	MaxIDLE           string `env:"REDIS_MAX_IDLE,required"`
	IDLETimeout       string `env:"REDIS_IDLE_TIMEOUT,required"`
	SessionTTL        string `env:"SESSION_TTL,required"`
}

type redisConfig struct {
	row redisEnvConfig
}

func NewRedisConfig() (*redisConfig, error) {
	var row redisEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &redisConfig{row: row}, nil
}
