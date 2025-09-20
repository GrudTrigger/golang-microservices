package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPORT = 50051
)

type InventoryService struct {
	inventoryV1.UnimplementedInventoryServiceServer
	mu    sync.RWMutex
	parts map[string]*inventoryV1.Part
}

func (s *InventoryService) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.parts[req.Uuid]
	if !ok {
		return nil, errors.New("NotFound")
	}
	return &inventoryV1.GetPartResponse{Part: p}, nil
}

func (s *InventoryService) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	return nil, nil
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

	service := &InventoryService{
		parts: make(map[string]*inventoryV1.Part),
	}

	inventoryV1.RegisterInventoryServiceServer(s, service)

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
