package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	inventoryApi "github.com/rocket-crm/inventory/internal/api/inventory/v1"
	"github.com/rocket-crm/inventory/internal/config"
	inventoryRepository "github.com/rocket-crm/inventory/internal/repository/inventory"
	inventoryService "github.com/rocket-crm/inventory/internal/service/inventory"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const configPath = "../deploy/compose/inventory/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}

	defer func() {
		cerr := client.Disconnect(ctx)
		if cerr != nil {
			log.Printf("failed to disconnect: %v\n", cerr)
		}
	}()

	db := client.Database(config.AppConfig().Mongo.DatabaseName())

	lis, err := net.Listen("tcp", config.AppConfig().InventoryGRPC.Address())
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	s := grpc.NewServer()

	repo := inventoryRepository.NewRepository(db)
	service := inventoryService.NewService(repo)
	api := inventoryApi.NewAPI(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)

	// Включаем рефлексию для отладки
	reflection.Register(s)
	go func() {
		log.Printf("🚀 gRPC server listening on %s\n", config.AppConfig().InventoryGRPC.Address())
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
