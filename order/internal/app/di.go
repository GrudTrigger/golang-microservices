package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rocker-crm/platform/pkg/closer"
	wrappedKafka "github.com/rocker-crm/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/rocker-crm/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/rocker-crm/platform/pkg/kafka/producer"
	"github.com/rocker-crm/platform/pkg/logger"
	kafkaMiddleware "github.com/rocker-crm/platform/pkg/middleware/kafka"
	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	orderAPI "github.com/rocket-crm/order/internal/api/order/v1"
	grpcClient "github.com/rocket-crm/order/internal/client/grpc"
	grpcInventory "github.com/rocket-crm/order/internal/client/grpc/inventory/v1"
	grpcPayment "github.com/rocket-crm/order/internal/client/grpc/payment/v1"
	"github.com/rocket-crm/order/internal/config"
	"github.com/rocket-crm/order/internal/migrator"
	"github.com/rocket-crm/order/internal/repository"
	"github.com/rocket-crm/order/internal/repository/order"
	"github.com/rocket-crm/order/internal/service"
	shipConsumer "github.com/rocket-crm/order/internal/service/consumer/ship_assembly"
	orderService "github.com/rocket-crm/order/internal/service/order"
	orderPaid "github.com/rocket-crm/order/internal/service/producer/order_paid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	orderV1API orderAPI.API

	ordersService    service.OrderService
	ordersRepository repository.OrderRepository

	inventoryClient grpcClient.InventoryClient
	paymentClient   grpcClient.PaymentClient
	authClient      authV1.AuthServiceClient

	syncProducer      sarama.SyncProducer
	orderPaidProducer wrappedKafka.Producer
	orderPaidService  service.ProducerService

	consumerGroup                 sarama.ConsumerGroup
	shipAssembledConsumer         service.ConsumerService
	shipAssembledRecorderConsumer wrappedKafka.Consumer

	postgresDb *pgx.Conn
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1API(ctx context.Context) orderAPI.API {
	if d.orderV1API == nil {
		d.orderV1API = orderAPI.NewAPI(d.Service(ctx))
	}
	return d.orderV1API
}

func (d *diContainer) Service(ctx context.Context) service.OrderService {
	if d.ordersService == nil {
		d.ordersService = orderService.NewService(d.Repository(ctx), d.InventoryClient(ctx), d.PaymentClient(ctx), d.OrderPaidService())
	}
	return d.ordersService
}

func (d *diContainer) AuthClient() authV1.AuthServiceClient {
	if d.authClient == nil {
		connIam, err := grpc.NewClient(config.AppConfig().OrderHttp.IamClientAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect: %w", err))
		}
		closer.AddNamed("Auth Client", func(ctx context.Context) error {
			return connIam.Close()
		})
		authClient := authV1.NewAuthServiceClient(connIam)
		d.authClient = authClient
	}
	return d.authClient
}

func (d *diContainer) InventoryClient(_ context.Context) grpcClient.InventoryClient {
	if d.inventoryClient == nil {
		connInventory, err := grpc.NewClient(config.AppConfig().OrderHttp.InventoryClientAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect: %w", err))
		}
		closer.AddNamed("Inventory Client", func(ctx context.Context) error {
			return connInventory.Close()
		})
		inventoryClient := inventoryV1.NewInventoryServiceClient(connInventory)
		grpcClientInventory := grpcInventory.NewClient(inventoryClient)
		d.inventoryClient = grpcClientInventory
	}
	return d.inventoryClient
}

func (d *diContainer) PaymentClient(_ context.Context) grpcClient.PaymentClient {
	if d.paymentClient == nil {
		connPayment, err := grpc.NewClient(config.AppConfig().OrderHttp.PaymentClientAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect: %w", err))
		}
		closer.AddNamed("Payment Client", func(ctx context.Context) error {
			return connPayment.Close()
		})
		paymentClient := paymentV1.NewPaymentServiceClient(connPayment)
		grpcClientPayment := grpcPayment.NewClient(paymentClient)
		d.paymentClient = grpcClientPayment
	}
	return d.paymentClient
}

func (d *diContainer) Repository(ctx context.Context) repository.OrderRepository {
	if d.ordersRepository == nil {
		d.ordersRepository = order.NewRepository(d.PostgresDb(ctx))
	}
	return d.ordersRepository
}

func (d *diContainer) PostgresDb(ctx context.Context) *pgx.Conn {
	if d.postgresDb == nil {
		conn, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}

		err = conn.Ping(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}

		migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), config.AppConfig().Postgres.MigrationDir())

		err = migratorRunner.Up()
		if err != nil {
			panic(fmt.Errorf("failed migrations up: %w", err))
		}

		closer.AddNamed("Postgres connect", func(ctx context.Context) error {
			return conn.Close(ctx)
		})

		d.postgresDb = conn
	}
	return d.postgresDb
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().Producer.Config(),
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

func (d *diContainer) OrderPaidProducer() wrappedKafka.Producer {
	if d.orderPaidProducer == nil {
		d.orderPaidProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().Producer.Topic(),
			logger.Logger(),
		)
	}
	return d.orderPaidProducer
}

func (d *diContainer) OrderPaidService() service.ProducerService {
	if d.orderPaidService == nil {
		d.orderPaidService = orderPaid.NewService(d.OrderPaidProducer())
	}
	return d.orderPaidService
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().Consumer.GroupID(),
			config.AppConfig().Consumer.Config(),
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

func (d *diContainer) ShipAssembledRecorderConsumer() wrappedKafka.Consumer {
	if d.shipAssembledRecorderConsumer == nil {
		d.shipAssembledRecorderConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{config.AppConfig().Consumer.Topic()},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}
	return d.shipAssembledRecorderConsumer
}

func (d *diContainer) ShipAssembledConsumer(ctx context.Context) service.ConsumerService {
	if d.shipAssembledConsumer == nil {
		d.shipAssembledConsumer = shipConsumer.NewService(d.ShipAssembledRecorderConsumer(), d.Repository(ctx))
	}
	return d.shipAssembledConsumer
}
