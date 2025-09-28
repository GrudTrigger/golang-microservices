package order

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/converter"
	"github.com/rocket-crm/order/internal/model"
)

func (s *ServiceSuite) TestCreateSuccess() {
	uuidParts := []string{uuid.NewString(), uuid.NewString()}
	filter := model.PartsFilter{Uuids: uuidParts}
	parts := []model.Part{
		{
			Uuid:          uuid.NewString(),
			Name:          "Engine Core",
			Description:   "High-performance spaceship engine core",
			Price:         1200,
			StockQuantity: 5,
			Category:      1,
			Dimensions: &model.Dimensions{
				Length: 2.5,
				Width:  1.2,
				Height: 1.8,
				Weight: 350.0,
			},
			Manufacturer: &model.Manufacturer{
				Name:    "SpaceTech Inc",
				Country: "USA",
				Website: "https://spacetech.example.com",
			},
			Tags: []string{"engine", "spaceship", "core"},
			Metadata: map[string]any{
				"serial": "SN-12345",
				"batch":  int64(42),
			},
			CreatedAt: converter.Ptr(time.Now()),
			UpdatedAt: converter.Ptr(time.Now()),
		},
		{
			Uuid:          uuid.NewString(),
			Name:          "Warp Drive",
			Description:   "Advanced warp drive for interstellar travel",
			Price:         2500,
			StockQuantity: 3,
			Category:      2,
			Dimensions: &model.Dimensions{
				Length: 3.0,
				Width:  1.5,
				Height: 2.0,
				Weight: 420.0,
			},
			Manufacturer: &model.Manufacturer{
				Name:    "Galactic Motors",
				Country: "UK",
				Website: "https://galacticmotors.example.com",
			},
			Tags: []string{"warp", "engine", "spaceship"},
			Metadata: map[string]any{
				"serial": "SN-67890",
				"batch":  int64(84),
			},
			CreatedAt: converter.Ptr(time.Now()),
			UpdatedAt: converter.Ptr(time.Now()),
		},
	}
	ctx := context.Background()
	s.inventoryClient.On("ListParts", ctx, filter).Return(parts, nil)

	createModel := model.CreateOrder{
		UserUUID:  uuid.NewString(),
		PartUuids: uuidParts,
	}

	var totalPrice float32 = 3700

	order := model.Order{
		OrderUUID:  uuid.NewString(),
		TotalPrice: totalPrice,
		PartUuids:  uuidParts,
		UserUUID:   uuid.NewString(),
		Status:     PAID,
	}
	s.orderRepository.On("Create", createModel, totalPrice).Return(order)

	resp, err := s.service.CreateOrder(ctx, createModel)
	s.Require().NoError(err)
	s.Require().Equal(totalPrice, resp.TotalPrice)
	s.Require().Equal(order.OrderUUID, resp.UUID)
}

func (s *ServiceSuite) TestCreateError() {
	uuidParts := []string{uuid.NewString(), uuid.NewString()}
	filter := model.PartsFilter{Uuids: uuidParts}
	errorClient := errors.New("найдены не все запчасти, проверьте uuid деталей")
	ctx := context.Background()
	s.inventoryClient.On("ListParts", ctx, filter).Return([]model.Part{}, errorClient)

	createModel := model.CreateOrder{
		UserUUID:  uuid.NewString(),
		PartUuids: uuidParts,
	}

	resp, err := s.service.CreateOrder(ctx, createModel)
	s.Require().ErrorIs(err, errorClient)
	s.Require().Empty(resp)
}
