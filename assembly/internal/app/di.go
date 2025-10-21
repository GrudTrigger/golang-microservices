package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rocker-crm/assembly/internal/config"
	"github.com/rocker-crm/assembly/internal/service"
	orderConsumer "github.com/rocker-crm/assembly/internal/service/consumer/order_consumer"
	"github.com/rocker-crm/platform/pkg/closer"
	wrappedKafka "github.com/rocker-crm/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/rocker-crm/platform/pkg/kafka/consumer"
	"github.com/rocker-crm/platform/pkg/logger"
	kafkaMiddleware "github.com/rocker-crm/platform/pkg/middleware/kafka"
)

type diContainer struct {
	orderPaidConsumer     service.ConsumerService
	consumerGroup         sarama.ConsumerGroup
	syncProducer          sarama.SyncProducer
	orderRecorderConsumer wrappedKafka.Consumer
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderPaidConsumer() service.ConsumerService {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = orderConsumer.NewConsumerService(d.OrderRecorderConsumer())
	}
	return d.orderPaidConsumer
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})

		d.consumerGroup = consumerGroup
	}
	return d.consumerGroup
}

func (d *diContainer) OrderRecorderConsumer() wrappedKafka.Consumer {
	if d.orderRecorderConsumer == nil {
		d.orderRecorderConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}
	return d.orderRecorderConsumer
}
