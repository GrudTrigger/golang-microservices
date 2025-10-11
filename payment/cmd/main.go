package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	paymentApi "github.com/rocket-crm/payment/internal/api/payment/v1"
	"github.com/rocket-crm/payment/internal/config"
	paymentRepository "github.com/rocket-crm/payment/internal/repository/payment"
	paymentService "github.com/rocket-crm/payment/internal/service/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const configPath = "../deploy/compose/payment/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	lis, err := net.Listen("tcp", config.AppConfig().PaymentGRPC.Address())
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

	repo := paymentRepository.NewRepository()
	service := paymentService.NewService(repo)
	api := paymentApi.NewAPI(service)
	paymentV1.RegisterPaymentServiceServer(s, api)

	// Включаем рефлексию для отладки
	reflection.Register(s)
	go func() {
		log.Printf("🚀 gRPC server listening on %s\n", config.AppConfig().PaymentGRPC.Address())
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
