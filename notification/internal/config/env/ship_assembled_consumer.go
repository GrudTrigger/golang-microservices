package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type shipAssembledConsumerEnvConfig struct {
	Topic   string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
	GroupID string `env:"ORDER_ASSEMBLED_CONSUMER_GROUP_ID,required"`
}

type shipAssembledConsumerConfig struct {
	row shipAssembledConsumerEnvConfig
}

func NewShipAssembledConsumerConfig() (*shipAssembledConsumerConfig, error) {
	var row shipAssembledConsumerEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &shipAssembledConsumerConfig{row: row}, nil
}

func (s *shipAssembledConsumerConfig) Topic() string {
	return s.row.Topic
}

func (s *shipAssembledConsumerConfig) GroupID() string {
	return s.row.GroupID
}

func (s *shipAssembledConsumerConfig) Config() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V4_0_0_0
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	return cfg
}
