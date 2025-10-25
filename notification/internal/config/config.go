package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rocker-crm/notifacation/internal/config/env"
)

var appConfig *config

type config struct {
	Logger                LoggerConfig
	Kafka                 KafkaConfig
	Bot                   BotConfig
	OrderPaidConsumer     OrderPaidConsumerConfig
	ShipAssembledConsumer ShipAssembledConsumerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderPaidConsumerCfg, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	shipAssembledConsumerCfg, err := env.NewShipAssembledConsumerConfig()
	if err != nil {
		return err
	}

	bot, err := env.NewBotConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                loggerCfg,
		Kafka:                 kafkaCfg,
		OrderPaidConsumer:     orderPaidConsumerCfg,
		ShipAssembledConsumer: shipAssembledConsumerCfg,
		Bot:                   bot,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
