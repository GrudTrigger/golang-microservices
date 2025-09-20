package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPORT = 50052
)

type PaymentService struct {
	paymentV1.UnimplementedPaymentServiceServer
}

func (s *PaymentService) PayOrder(context.Context, *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	tranUuid := uuid.NewString()
	log.Printf("Оплата прошла успешно, transaction_uuid: <%s>\n", tranUuid)
	return &paymentV1.PayOrderResponse{TransactionUuid: tranUuid}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPORT))
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

	service := &PaymentService{}

	paymentV1.RegisterPaymentServiceServer(s, service)

	// Включаем рефлексию для отладки
	reflection.Register(s)
	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPORT)
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
