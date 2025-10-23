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
	Kafka     KafkaConfig
	Producer  OrderPaidProducerConfig
	Consumer  ShipAssembledConsumerConfig
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

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	producer, err := env.NewOrderPaidProducerConfig()
	if err != nil {
		return err
	}

	consumer, err := env.NewShipAssembledConsumerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:    logger,
		Postgres:  postgres,
		OrderHttp: orderHttp,
		Kafka:     kafka,
		Producer:  producer,
		Consumer:  consumer,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
