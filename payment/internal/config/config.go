package config

import (
	"github.com/joho/godotenv"
	"github.com/rocket-crm/payment/internal/config/env"
)

var appConfig *config

type config struct {
	Logger      LoggerConfig
	PaymentGRPC PaymentGRPCConfig
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

	paymentGRPC, err := env.NewPaymentConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:      logger,
		PaymentGRPC: paymentGRPC,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
