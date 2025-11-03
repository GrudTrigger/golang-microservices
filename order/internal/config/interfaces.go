package config

import "github.com/IBM/sarama"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type PostgresConfig interface {
	URI() string
	MigrationDir() string
}

type OrderHttpConfig interface {
	InventoryClientAddress() string
	PaymentClientAddress() string
	IamClientAddress() string
	Address() string
	TimeOut() string
}

type KafkaConfig interface {
	Brokers() []string
}

type OrderPaidProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type ShipAssembledConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}
