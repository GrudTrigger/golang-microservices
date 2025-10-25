package env

import "github.com/caarlos0/env/v11"

type kafkaEnvConfig struct {
	Brokers []string `env:"KAFKA_BROKERS,required"`
}

type kafkaConfig struct {
	row kafkaEnvConfig
}

func NewKafkaConfig() (*kafkaConfig, error) {
	var row kafkaEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &kafkaConfig{row: row}, nil
}

func (k *kafkaConfig) Brokers() []string {
	return k.row.Brokers
}
