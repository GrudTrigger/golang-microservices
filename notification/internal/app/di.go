package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/go-telegram/bot"
	"github.com/rocker-crm/notifacation/internal/client"
	"github.com/rocker-crm/notifacation/internal/client/http/telegram"
	"github.com/rocker-crm/notifacation/internal/config"
	services "github.com/rocker-crm/notifacation/internal/service"
	"github.com/rocker-crm/notifacation/internal/service/consumers/order_paid_consumer"
	"github.com/rocker-crm/notifacation/internal/service/consumers/ship_assembled_consumer"
	tgService "github.com/rocker-crm/notifacation/internal/service/telegram"
	"github.com/rocker-crm/platform/pkg/closer"
	wrappedKafka "github.com/rocker-crm/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/rocker-crm/platform/pkg/kafka/consumer"
	"github.com/rocker-crm/platform/pkg/logger"
	kafkaMiddleware "github.com/rocker-crm/platform/pkg/middleware/kafka"
)

type diContainer struct {
	orderRecorderConsumer wrappedKafka.Consumer
	shipRecorderConsumer  wrappedKafka.Consumer

	orderConsumerGroup sarama.ConsumerGroup
	shipConsumerGroup  sarama.ConsumerGroup

	orderPaidConsumerService     services.ConsumerService
	shipAssembledConsumerService services.ConsumerService

	telegramBot     *bot.Bot
	telegramClient  client.TelegramClient
	telegramService services.TelegramService
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderPaidConsumerService() services.ConsumerService {
	if d.orderPaidConsumerService == nil {
		d.orderPaidConsumerService = order_paid_consumer.NewOrderPaidConsumerService(d.OrderRecorderConsumer(), d.TelegramService())
	}
	return d.orderPaidConsumerService
}

func (d *diContainer) ShipAssembledConsumerService() services.ConsumerService {
	if d.shipAssembledConsumerService == nil {
		d.shipAssembledConsumerService = ship_assembled_consumer.NewShipAssembledConsumerService(d.ShipRecorderConsumer(), d.TelegramService())
	}
	return d.shipAssembledConsumerService
}

func (d *diContainer) OrderRecorderConsumer() wrappedKafka.Consumer {
	if d.orderRecorderConsumer == nil {
		d.orderRecorderConsumer = wrappedKafkaConsumer.NewConsumer(
			d.OrderConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}
	return d.orderRecorderConsumer
}

func (d *diContainer) OrderConsumerGroup() sarama.ConsumerGroup {
	if d.orderConsumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.orderConsumerGroup.Close()
		})
		d.orderConsumerGroup = consumerGroup
	}
	return d.orderConsumerGroup
}

func (d *diContainer) ShipRecorderConsumer() wrappedKafka.Consumer {
	if d.shipRecorderConsumer == nil {
		d.shipRecorderConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ShipConsumerGroup(),
			[]string{
				config.AppConfig().ShipAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}
	return d.shipRecorderConsumer
}

func (d *diContainer) ShipConsumerGroup() sarama.ConsumerGroup {
	if d.shipConsumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().ShipAssembledConsumer.GroupID(),
			config.AppConfig().ShipAssembledConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.orderConsumerGroup.Close()
		})
		d.shipConsumerGroup = consumerGroup
	}
	return d.shipConsumerGroup
}

func (d *diContainer) TelegramBot() *bot.Bot {
	if d.telegramBot == nil {
		b, err := bot.New(config.AppConfig().Bot.Token())
		if err != nil {
			panic(fmt.Sprintf("failed to create telegram bot: %s\n", err.Error()))
		}

		d.telegramBot = b
	}
	return d.telegramBot
}

func (d *diContainer) TelegramClient() client.TelegramClient {
	if d.telegramClient == nil {
		d.telegramClient = telegram.NewClient(d.TelegramBot())
	}
	return d.telegramClient
}

func (d *diContainer) TelegramService() services.TelegramService {
	if d.telegramService == nil {
		d.telegramService = tgService.NewService(d.TelegramClient())
	}
	return d.telegramService
}
