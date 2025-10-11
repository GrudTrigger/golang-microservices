package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rocker-crm/platform/pkg/closer"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	orderAPI "github.com/rocket-crm/order/internal/api/order/v1"
	grpcClient "github.com/rocket-crm/order/internal/client/grpc"
	grpcInventory "github.com/rocket-crm/order/internal/client/grpc/inventory/v1"
	grpcPayment "github.com/rocket-crm/order/internal/client/grpc/payment/v1"
	"github.com/rocket-crm/order/internal/migrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rocket-crm/order/internal/config"
	"github.com/rocket-crm/order/internal/repository"
	"github.com/rocket-crm/order/internal/repository/order"
	"github.com/rocket-crm/order/internal/service"
	orderService "github.com/rocket-crm/order/internal/service/order"
)

type diContainer struct {
	orderV1API orderAPI.API

	ordersService    service.OrderService
	ordersRepository repository.OrderRepository

	inventoryClient grpcClient.InventoryClient
	paymentClient   grpcClient.PaymentClient

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
		d.ordersService = orderService.NewService(d.Repository(ctx), d.InventoryClient(ctx), d.PaymentClient(ctx))
	}
	return d.ordersService
}

func (d *diContainer) InventoryClient(_ context.Context) grpcClient.InventoryClient {
	if d.inventoryClient == nil {
		connInventory, err := grpc.NewClient(config.AppConfig().OrderHttp.InventoryClientAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect: %w\n", err))
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
			panic(fmt.Errorf("failed to connect: %w\n", err))
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
			panic(fmt.Errorf("failed to connect to database: %w\n", err))
		}

		err = conn.Ping(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w\n", err))
		}

		migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), config.AppConfig().Postgres.MigrationDir())

		err = migratorRunner.Up()
		if err != nil {
			panic(fmt.Errorf("failed migrations up: %w\n", err))
		}

		closer.AddNamed("Postgres connect", func(ctx context.Context) error {
			return conn.Close(ctx)
		})
		
		d.postgresDb = conn
	}
	return d.postgresDb
}
