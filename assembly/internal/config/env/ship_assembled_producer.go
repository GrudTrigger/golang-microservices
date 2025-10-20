package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type shipAssembledProducerEnvConfig struct {
	TopicName string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
}

type shipAssembledProducerConfig struct {
	row shipAssembledProducerEnvConfig
}

func NewShipAssembledProducerConfig() (*shipAssembledProducerConfig, error) {
	var row shipAssembledProducerEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &shipAssembledProducerConfig{row: row}, nil
}

func (s *shipAssembledProducerConfig) Topic() string {
	return s.row.TopicName
}

func (s *shipAssembledProducerConfig) Config() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V4_0_0_0
	cfg.Producer.Return.Successes = true
	return cfg
}
