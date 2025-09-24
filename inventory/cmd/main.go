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

	"github.com/google/uuid"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	var res *inventoryV1.Part
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.parts {
		if p.Uuid == req.Uuid {
			res = p
		}
	}
	if res == nil {
		return nil, errors.New("NotFound")
	}
	return &inventoryV1.GetPartResponse{Part: res}, nil
}

func (s *InventoryService) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	var res []*inventoryV1.Part
	
	if len(req.Filter.Uuids) == 0 && len(req.Filter.ManufacturerCountries) == 0 && len(req.Filter.Names) == 0 && len(req.Filter.Tags) == 0 && len(req.Filter.Categories) == 0 {
		for _, p := range s.parts {
			res = append(res, p)
		}
		return &inventoryV1.ListPartsResponse{Parts: res}, nil
	}

	if len(req.Filter.Uuids) > 0 {
		for _, v := range req.Filter.Uuids {
			for _, p := range s.parts {
				if p.Uuid == v {
					res = append(res, p)
				}
			}
		}
	}

	if len(req.Filter.Names) > 0 {
		for _, v := range req.Filter.Names {
			for _, p := range s.parts {
				if p.Name == v {
					res = append(res, p)
				}
			}
		}
	}
	if len(req.Filter.Categories) > 0 {
		for _, v := range req.Filter.Categories {
			for _, p := range s.parts {
				if v == p.Category {
					res = append(res, p)
				}
			}
		}
	}
	if len(req.Filter.ManufacturerCountries) > 0 {
		for _, v := range req.Filter.ManufacturerCountries {
			for _, p := range s.parts {
				if v == p.Manufacturer.Country {
					res = append(res, p)
				}
			}
		}
	}
	if len(req.Filter.Tags) > 0 {
		for _, v := range req.Filter.Tags {
			for _, p := range s.parts {
				for _, t := range p.Tags {
					if v == t {
						res = append(res, p)
					}
				}
			}
		}
	}
	return &inventoryV1.ListPartsResponse{Parts: res}, nil
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
		parts: map[string]*inventoryV1.Part{
			"part-1": {
				Uuid:          uuid.NewString(),
				Name:          "Engine Core",
				Description:   "High-performance spaceship engine core",
				Price:         1200.50,
				StockQuantity: 5,
				Category:      inventoryV1.Category_CATEGORY_ENGINE,
				Dimensions: &inventoryV1.Dimensions{
					Length: 2.5,
					Width:  1.2,
					Height: 1.8,
					Weight: 350.0,
				},
				Manufacturer: &inventoryV1.Manufacturer{
					Name:    "SpaceTech Inc",
					Country: "USA",
					Website: "https://spacetech.example.com",
				},
				Tags: []string{"engine", "spaceship", "core"},
				Metadata: map[string]*inventoryV1.Value{
					"serial": {Kind: &inventoryV1.Value_StringValue{StringValue: "SN-12345"}},
					"batch":  {Kind: &inventoryV1.Value_Int64Value{Int64Value: 42}},
				},
				CreatedAt: timestamppb.Now(),
				UpdatedAt: timestamppb.Now(),
			},
			"part-2": {
				Uuid:          uuid.NewString(),
				Name:          "Fuel Pump",
				Description:   "Reliable fuel pump for interstellar flights",
				Price:         300.75,
				StockQuantity: 12,
				Category:      inventoryV1.Category_CATEGORY_FUEL,
				Dimensions: &inventoryV1.Dimensions{
					Length: 0.8,
					Width:  0.4,
					Height: 0.6,
					Weight: 25.0,
				},
				Manufacturer: &inventoryV1.Manufacturer{
					Name:    "Galaxy Supplies",
					Country: "Germany",
					Website: "https://galaxy-supplies.example.com",
				},
				Tags: []string{"fuel", "pump"},
				Metadata: map[string]*inventoryV1.Value{
					"serial": {Kind: &inventoryV1.Value_StringValue{StringValue: "FP-00987"}},
				},
				CreatedAt: timestamppb.Now(),
				UpdatedAt: timestamppb.Now(),
			},
			"part-3": {
				Uuid:          uuid.NewString(),
				Name:          "Porthole Glass",
				Description:   "Reinforced glass for spaceship portholes",
				Price:         150.00,
				StockQuantity: 50,
				Category:      inventoryV1.Category_CATEGORY_PORTHOLE,
				Dimensions: &inventoryV1.Dimensions{
					Length: 1.0,
					Width:  1.0,
					Height: 0.02,
					Weight: 10.0,
				},
				Manufacturer: &inventoryV1.Manufacturer{
					Name:    "CosmoGlass",
					Country: "Japan",
					Website: "https://cosmoglass.example.com",
				},
				Tags: []string{"glass", "window"},
				Metadata: map[string]*inventoryV1.Value{
					"heat_resistance": {Kind: &inventoryV1.Value_DoubleValue{DoubleValue: 1250.5}},
				},
				CreatedAt: timestamppb.Now(),
				UpdatedAt: timestamppb.Now(),
			},
		},
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
