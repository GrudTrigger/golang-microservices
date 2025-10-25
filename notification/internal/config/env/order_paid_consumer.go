package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type orderPaidConsumerEnvConfig struct {
	Topic   string `env:"ORDER_PAID_TOPIC_NAME,required"`
	GroupID string `env:"ORDER_PAID_CONSUMER_GROUP_ID,required"`
}

type orderPaidConsumerConfig struct {
	row orderPaidConsumerEnvConfig
}

func NewOrderPaidConsumerConfig() (*orderPaidConsumerConfig, error) {
	var row orderPaidConsumerEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &orderPaidConsumerConfig{row: row}, nil
}

func (o *orderPaidConsumerConfig) Topic() string {
	return o.row.Topic
}

func (o *orderPaidConsumerConfig) GroupID() string {
	return o.row.GroupID
}

func (o *orderPaidConsumerConfig) Config() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V4_0_0_0
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	return cfg
}
