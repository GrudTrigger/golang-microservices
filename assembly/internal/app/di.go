package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rocker-crm/assembly/internal/config"
	"github.com/rocker-crm/assembly/internal/service"
	orderConsumer "github.com/rocker-crm/assembly/internal/service/consumer/order_consumer"
	shipProducer "github.com/rocker-crm/assembly/internal/service/producer/ship_assembled"
	"github.com/rocker-crm/platform/pkg/closer"
	wrappedKafka "github.com/rocker-crm/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/rocker-crm/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/rocker-crm/platform/pkg/kafka/producer"
	"github.com/rocker-crm/platform/pkg/logger"
	kafkaMiddleware "github.com/rocker-crm/platform/pkg/middleware/kafka"
)

type diContainer struct {
	orderPaidConsumer     service.ConsumerService
	consumerGroup         sarama.ConsumerGroup
	orderRecorderConsumer wrappedKafka.Consumer
	syncProducer          sarama.SyncProducer
	shipAssembledProducer wrappedKafka.Producer
	shipAssembledService  service.ProducerService
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderPaidConsumer() service.ConsumerService {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = orderConsumer.NewConsumerService(d.OrderRecorderConsumer(), d.ShipAssembledService())
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

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().ShipAssembledProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}
	return d.syncProducer
}

func (d *diContainer) ShipAssembledProducer() wrappedKafka.Producer {
	if d.shipAssembledProducer == nil {
		d.shipAssembledProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().ShipAssembledProducer.Topic(),
			logger.Logger(),
		)
	}
	return d.shipAssembledProducer
}

func (d *diContainer) ShipAssembledService() service.ProducerService {
	if d.shipAssembledService == nil {
		d.shipAssembledService = shipProducer.NewService(d.ShipAssembledProducer())
	}
	return d.shipAssembledService
}
