package config

import (
	"github.com/joho/godotenv"
	"github.com/rocket-crm/order/internal/config/env"
)

var appConfig *config

type config struct {
	Logger    LoggerConfig
	OrderHttp OrderHttpConfig
	Postgres  PostgresConfig
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

	orderHttp, err := env.NewOrderHttpConfig()
	if err != nil {
		return err
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:    logger,
		Postgres:  postgres,
		OrderHttp: orderHttp,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
