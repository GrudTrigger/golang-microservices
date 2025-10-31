package config

import (
	"github.com/joho/godotenv"
	"github.com/rocket-crm/iam/internal/config/env"
)

var appConfig *config

type config struct {
	Logger   LoggerConfig
	Postgres PostgresConfig
	IamGRPC  IAMGRPCConfig
	Redis    RedisConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}

	logger, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	authGRPC, err := env.NewAuthGRPCConfig()
	if err != nil {
		return err
	}

	redis, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:   logger,
		Postgres: postgres,
		IamGRPC:  authGRPC,
		Redis:    redis,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
