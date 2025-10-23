package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type orderPaidProducerEnvConfig struct {
	TopicName string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
}

type orderPaidProducerConfig struct {
	row orderPaidProducerEnvConfig
}

func NewShipAssembledProducerConfig() (*orderPaidProducerConfig, error) {
	var row orderPaidProducerEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &orderPaidProducerConfig{row: row}, nil
}

func (s *orderPaidProducerConfig) Topic() string {
	return s.row.TopicName
}

func (s *orderPaidProducerConfig) Config() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V4_0_0_0
	cfg.Producer.Return.Successes = true
	return cfg
}
