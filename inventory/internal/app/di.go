package app

import (
	"context"
	"fmt"

	"github.com/rocker-crm/platform/pkg/closer"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	inventoryV1API "github.com/rocket-crm/inventory/internal/api/inventory/v1"
	"github.com/rocket-crm/inventory/internal/config"
	"github.com/rocket-crm/inventory/internal/repository"
	"github.com/rocket-crm/inventory/internal/repository/inventory"
	"github.com/rocket-crm/inventory/internal/service"
	inventoryService "github.com/rocket-crm/inventory/internal/service/inventory"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type diContainer struct {
	inventoryV1API      inventoryV1.InventoryServiceServer
	inventoryService    service.InventoryService
	inventoryRepository repository.InventoryRepository

	mongoDBClient *mongo.Client
	mongoDBHandle *mongo.Database
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) InventoryV1API(ctx context.Context) inventoryV1.InventoryServiceServer {
	if d.inventoryV1API == nil {
		d.inventoryV1API = inventoryV1API.NewAPI(d.Service(ctx))
	}
	return d.inventoryV1API
}

func (d *diContainer) Service(ctx context.Context) service.InventoryService {
	if d.inventoryService == nil {
		d.inventoryService = inventoryService.NewService(d.Repository(ctx))
	}
	return d.inventoryService
}

func (d *diContainer) Repository(ctx context.Context) repository.InventoryRepository {
	if d.inventoryRepository == nil {
		d.inventoryRepository = inventory.NewRepository(d.MongoDBHandle(ctx))
	}
	return d.inventoryRepository
}

func (d *diContainer) MongoClient(ctx context.Context) *mongo.Client {
	if d.mongoDBClient == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to MongoDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("failed to ping MongoDB: %v\n", err))
		}

		closer.AddNamed("Mongo DB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})
		d.mongoDBClient = client
	}
	return d.mongoDBClient
}

func (d *diContainer) MongoDBHandle(ctx context.Context) *mongo.Database {
	if d.mongoDBHandle == nil {
		d.mongoDBHandle = d.MongoClient(ctx).Database(config.AppConfig().Mongo.DatabaseName())
	}
	return d.mongoDBHandle
}
